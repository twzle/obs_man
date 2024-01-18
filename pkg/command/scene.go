package command

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"go.uber.org/zap"
)

func ProvideSceneCommands(obsManager ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(SetCurrentPreviewScene{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := SetCurrentPreviewScene{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(SetCurrentProgramScene{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := SetCurrentProgramScene{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
	}
}

/*----------------------------- Set Preview/Current --------------------------*/

var _ RunnableCommand = &SetCurrentProgramScene{}
var _ RunnableCommand = &SetCurrentPreviewScene{}

type SetCurrentProgramScene struct {
	ProgramSceneName string `hubman:"program_scene_name"`
}

func (s SetCurrentProgramScene) Run(p ObsProvider, _ *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return err
	}
	_, err = obsClient.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{
		SceneName: s.ProgramSceneName,
	})
	return err
}

func (s SetCurrentProgramScene) Code() string {
	return "SetCurrentProgramScene"
}

func (s SetCurrentProgramScene) Description() string {
	return "Sets current Program Scene"
}

type SetCurrentPreviewScene struct {
	PreviewSceneName string `hubman:"preview_scene_name"`
}

func (s SetCurrentPreviewScene) Run(p ObsProvider, _ *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return err
	}
	_, err = obsClient.Scenes.SetCurrentPreviewScene(&scenes.SetCurrentPreviewSceneParams{
		SceneName: s.PreviewSceneName,
	})
	return err
}

func (s SetCurrentPreviewScene) Code() string {
	return "SetCurrentPreviewScene"
}

func (s SetCurrentPreviewScene) Description() string {
	return "Sets current Preview Scene"
}

/*----------------------------- Set SceneItem --------------------------------*/
