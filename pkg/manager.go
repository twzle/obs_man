package obsman

import (
	"context"
	"errors"
	cmd "obs-man/pkg/command"
	osig "obs-man/pkg/signal"
	"reflect"
	"sync"
	"time"

	"git.miem.hse.ru/hubman/hubman-lib/core"
	"github.com/andreykaipov/goobs"
	obsevents "github.com/andreykaipov/goobs/api/events"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	_                         cmd.ObsProvider = &manager{}
	timeoutedDialer                           = &websocket.Dialer{HandshakeTimeout: 1500 * time.Millisecond}
	defaultHealthcheckMillis                  = 3000
	obsDisconnectedCheckLabel                 = "OBS_DISCONNECTED"
)

func connectObs(hostPort, password string) (*goobs.Client, error) {
	return goobs.New(
		hostPort,
		goobs.WithPassword(password),
		goobs.WithResponseTimeout(1000*time.Millisecond),
		goobs.WithDialer(timeoutedDialer),
	)
}

type manager struct {
	conf   ObsConf
	logger *zap.Logger
	client *goobs.Client
	mutex  sync.Mutex

	connected    bool
	listenCtx    context.Context
	cancelListen context.CancelFunc
	signals      chan<- core.Signal
	checkManager core.CheckRegistry
}

func (m *manager) listenEvents() {
	for {
		select {
		case ev := <-m.client.IncomingEvents:
			if ev != nil {
				m.logger.Debug(
					"Received event from obs",
					zap.Any("event", ev),
					zap.String("event_type", reflect.TypeOf(ev).String()),
				)
			}
			m.processObsEvent(ev)
		case <-m.listenCtx.Done():
			return
		}
	}
}

func (m *manager) processObsEvent(event interface{}) {
	switch e := event.(type) {
	case *obsevents.CurrentPreviewSceneChanged:
		sceneId, err := m.GetSceneId(e.SceneName)
		if err != nil {
			m.logger.Debug(
				"Preview scene changed signal processing error",
				zap.Error(err))
			return
		}
		m.signals <- &osig.CurrentPreviewSceneChangedById{
			SceneName: e.SceneName,
			SceneId:   sceneId,
		}

	case *obsevents.CurrentProgramSceneChanged:
		sceneId, err := m.GetSceneId(e.SceneName)
		if err != nil {
			m.logger.Debug(
				"Program scene changed signal processing error",
				zap.Error(err))
			return
		}
		m.signals <- &osig.CurrentProgramSceneChangedById{
			SceneName: e.SceneName,
			SceneId:   sceneId,
		}
	case *obsevents.InputMuteStateChanged:
		m.signals <- &osig.InputMuteStateChanged{
			InputName:  e.InputName,
			InputMuted: e.InputMuted,
		}
	case *obsevents.InputVolumeChanged:
		m.signals <- &osig.InputVolumeChanged{
			InputName:      e.InputName,
			InputVolumeMul: e.InputVolumeMul,
			InputVolumeDb:  e.InputVolumeDb,
		}
	case *obsevents.RecordStateChanged:
		m.signals <- &osig.RecordStateChanged{
			OutputActive: e.OutputActive,
			OutputState:  e.OutputState,
			OutputPath:   e.OutputPath,
		}
	case *obsevents.ReplayBufferStateChanged:
		m.signals <- &osig.ReplayBufferStateChanged{
			OutputActive: e.OutputActive,
			OutputState:  e.OutputState,
		}
	case *obsevents.ReplayBufferSaved:
		m.signals <- &osig.ReplayBufferSaved{
			SavedReplayPath: e.SavedReplayPath,
		}
	case *obsevents.SceneItemEnableStateChanged:
		m.signals <- &osig.SceneItemEnableStateChanged{
			SceneName:        e.SceneName,
			SceneItemId:      int(e.SceneItemId),
			SceneItemEnabled: e.SceneItemEnabled,
		}
	case *obsevents.ScreenshotSaved:
		m.signals <- &osig.ScreenshotSaved{
			SavedScreenshotPath: e.SavedScreenshotPath,
		}
	case *obsevents.StreamStateChanged:
		m.signals <- &osig.StreamStateChanged{
			OutputActive: e.OutputActive,
			OutputState:  e.OutputState,
		}
	case *obsevents.SceneTransitionStarted:
		m.signals <- &osig.SceneTransitionStarted{
			TransitionName: e.TransitionName,
		}
	case *obsevents.SceneTransitionEnded:
		m.signals <- &osig.SceneTransitionEnded{
			TransitionName: e.TransitionName,
		}
	case *obsevents.SceneTransitionVideoEnded:
		m.signals <- &osig.SceneTransitionVideoEnded{
			TransitionName: e.TransitionName,
		}
	case *obsevents.StudioModeStateChanged:
		m.signals <- &osig.StudioModeStateChanged{
			StudioModeEnabled: e.StudioModeEnabled,
		}
	case *obsevents.VirtualcamStateChanged:
		m.signals <- &osig.VirtualCamStateChanged{
			OutputActive: e.OutputActive,
			OutputState:  e.OutputState,
		}
	case nil, error, *obsevents.ExitStarted:
		m.Close()
	}
}

func (m *manager) GetSignals() chan<- core.Signal {
	return m.signals
}

func (m *manager) GetSceneId(sceneName string) (int, error) {
	obsClient, err := m.Provide()
	if err != nil {
		return -1, err
	}

	sceneListResponse, _ := obsClient.Scenes.GetSceneList()
	for _, scene := range sceneListResponse.Scenes {
		if scene.SceneName == sceneName {
			return len(sceneListResponse.Scenes) - scene.SceneIndex, nil
		}
	}
	return -1, errors.New("can't find scene with given name")
}

func (m *manager) Provide() (*goobs.Client, error) {
	if m.client == nil || !m.connected {
		return nil, errors.New("no opened obs connection")
	}
	return m.client, nil
}

func (m *manager) Close() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.connected && m.client != nil {
		m.cancelListen()
		_ = m.client.Disconnect()
	}
	m.connected = false

	m.checkManager.RegisterFail(core.NewCheck(
		obsDisconnectedCheckLabel,
		"obs was disconnected",
	))
	return nil
}

func (m *manager) UpdateConn(c ObsConf) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	ctxlog := m.logger.With(zap.Any("config", c))
	ctxlog.Debug("Updating obs connection")
	if c == m.conf && m.connected {
		ctxlog.Debug("Already connected")
		return nil
	}

	m.cancelListen()
	if m.connected {
		err := m.client.Disconnect()
		if err != nil {
			ctxlog.Error("Failed to disconnect from previous obs client", zap.Error(err))
		}
	}

	m.connected = false
	m.checkManager.RegisterFail(core.NewCheck(
		obsDisconnectedCheckLabel,
		"obs was disconnected",
	))
	m.conf = c

	client, err := connectObs(c.HostPort, c.Password)
	if err != nil {
		ctxlog.Error("Failed to connect to obs", zap.Error(err))
		return err
	}

	m.connected = true
	m.checkManager.RegisterSuccess(core.NewCheck(
		obsDisconnectedCheckLabel,
		"obs is connected",
	))

	m.client = client
	m.listenCtx, m.cancelListen = context.WithCancel(context.Background())

	ctxlog.Debug("Successfully updated obs connection")
	go m.listenEvents()

	return nil
}

func NewManager(c ObsConf, logger *zap.Logger, signalsCh chan<- core.Signal, checkManager core.CheckRegistry) (*manager, error) {
	connected := false
	client, err := connectObs(c.HostPort, c.Password)
	if err == nil {
		connected = true
		check := core.NewCheck(
			obsDisconnectedCheckLabel,
			"obs is connected",
		)
		checkManager.RegisterSuccess(check)
	}
	m := &manager{
		client:       client,
		logger:       logger.Named("OBSManager"),
		signals:      signalsCh,
		conf:         c,
		connected:    connected,
		checkManager: checkManager,
	}

	m.listenCtx, m.cancelListen = context.WithCancel(context.Background())
	if m.client == nil || !m.connected {
		return m, err
	}

	go m.listenEvents()

	return m, nil
}

func (m *manager) HealthCheck(c ObsConf, shutdown <-chan bool) {
	if c.HealthCheckInterval == 0 {
		c.HealthCheckInterval = defaultHealthcheckMillis
	}
	ticker := time.NewTicker(time.Duration(c.HealthCheckInterval) * time.Millisecond)
	disconnected := false

	for {
		select {
		case <-shutdown:
			m.Close()
			return
		case <-ticker.C:
			m.mutex.Lock()
			if m.client == nil || !m.connected {
				disconnected = true
			} else if m.client != nil && m.connected {
				disconnected = false
				if _, err := m.client.General.GetStats(); err != nil {
					m.connected = false
					m.client.Disconnect()
					m.checkManager.RegisterFail(core.NewCheck(
						obsDisconnectedCheckLabel,
						"obs was disconnected",
					))
					disconnected = true
				}
			}

			m.mutex.Unlock()

			if disconnected {

				m.mutex.Lock()
				curConf := m.conf
				m.mutex.Unlock()

				m.UpdateConn(curConf)
			}
		}
	}
}
