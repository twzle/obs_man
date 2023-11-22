package obsman

import (
	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"go.uber.org/zap"
	"io"
)

type Manager interface {
	io.Closer
	DoSetScene(SetScene)
	DoStartRecord(StartRecord)
	DoStopRecord(StopRecord)
}

type manager struct {
	logger *zap.Logger
	client *goobs.Client
}

func (m *manager) Close() error {
	_ = m.client.Disconnect()
	return nil
}

func (m *manager) DoSetScene(scene SetScene) {
	_, _ = m.client.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{SceneName: scene.SceneName})
}

func (m *manager) DoStartRecord(cmd StartRecord) {
	m.client.Record.StartRecord()
}

func (m *manager) DoStopRecord(cmd StopRecord) {
	m.client.Record.StopRecord()
}

func NewManager(c ObsConf, logger *zap.Logger) (*manager, error) {
	client, err := goobs.New(c.HostPort, goobs.WithPassword(c.Password))
	return &manager{client: client, logger: logger.Named("OBSManager")}, err
}
