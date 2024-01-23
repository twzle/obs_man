package obsman

import (
	"context"
	"errors"
	cmd "obs-man/pkg/command"
	osig "obs-man/pkg/signal"
	"reflect"

	"git.miem.hse.ru/hubman/hubman-lib/core"
	"github.com/andreykaipov/goobs"
	obsevents "github.com/andreykaipov/goobs/api/events"
	"go.uber.org/zap"
)

var _ cmd.ObsProvider = &manager{}

type manager struct {
	conf   ObsConf
	logger *zap.Logger
	client *goobs.Client

	connected    bool
	listenCtx    context.Context
	cancelListen context.CancelFunc
	signals      chan<- core.Signal
}

func (m *manager) listenEvents() {
	for {
		select {
		case ev := <-m.client.IncomingEvents:
			m.logger.Debug(
				"Received event from obs",
				zap.Any("event", ev),
				zap.String("event_type", reflect.TypeOf(ev).String()),
			)
			m.processObsEvent(ev)
		case <-m.listenCtx.Done():
			return
		}
	}
}

func (m *manager) processObsEvent(event interface{}) {
	switch e := event.(type) {
	case *obsevents.CurrentPreviewSceneChanged:
		m.signals <- &osig.CurrentPreviewSceneChanged{
			SceneName: e.SceneName,
		}
	case *obsevents.CurrentProgramSceneChanged:
		m.signals <- &osig.CurrentProgramSceneChanged{
			SceneName: e.SceneName,
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
	case *obsevents.ExitStarted:
		m.Close()
	}
}

func (m *manager) Provide() (*goobs.Client, error) {
	if m.client == nil {
		return nil, errors.New("no opened obs connection")
	}
	return m.client, nil
}

func (m *manager) Close() error {
	m.connected = false
	m.cancelListen()
	_ = m.client.Disconnect()
	return nil
}

func (m *manager) UpdateConn(c ObsConf) error {
	ctxlog := m.logger.With(zap.Any("config", c))
	ctxlog.Debug("Updating obs connection")
	if c == m.conf && m.connected {
		ctxlog.Debug("Already connected")
		return nil
	}

	client, err := goobs.New(c.HostPort, goobs.WithPassword(c.Password))
	if err != nil {
		ctxlog.Error("Failed to connect to obs", zap.Error(err))
		return err
	}

	m.cancelListen()
	m.client = client
	m.listenCtx, m.cancelListen = context.WithCancel(context.Background())

	go m.listenEvents()
	m.connected = true

	return nil
}

func NewManager(c ObsConf, logger *zap.Logger, signalsCh chan<- core.Signal) (*manager, error) {
	client, err := goobs.New(c.HostPort, goobs.WithPassword(c.Password))
	if err != nil {
		return nil, err
	}
	m := &manager{
		client:    client,
		logger:    logger.Named("OBSManager"),
		signals:   signalsCh,
		conf:      c,
		connected: true,
	}

	m.listenCtx, m.cancelListen = context.WithCancel(context.Background())
	go m.listenEvents()

	return m, nil
}
