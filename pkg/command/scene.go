package command

import (
	"errors"

	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"go.uber.org/zap"
)

// Fucntion providing handlers for command to manage scene switch process with OBS
func ProvideSceneCommands(obsManager ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(SetCurrentPreviewScene{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := SetCurrentPreviewScene{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(SetCurrentProgramScene{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := SetCurrentProgramScene{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(SetCurrentPreviewSceneById{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := SetCurrentPreviewSceneById{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(SetCurrentProgramSceneById{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := SetCurrentProgramSceneById{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
	}
}

/*----------------------------- Set Preview/Current --------------------------*/

var _ RunnableCommand = &SetCurrentProgramScene{}
var _ RunnableCommand = &SetCurrentPreviewScene{}
var _ RunnableCommand = &SetCurrentProgramSceneById{}
var _ RunnableCommand = &SetCurrentPreviewSceneById{}

// Representation of set current program scene command
type SetCurrentProgramScene struct {
	ProgramSceneName string `hubman:"program_scene_name"`
}

// Function provides handler to execute command in OBS
func (s SetCurrentProgramScene) Run(p ObsProvider, l *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return err
	}
	l.Error("Error executing command", zap.Error(err))
	_, err = obsClient.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{
		SceneName: &s.ProgramSceneName,
	})
	if err != nil {
		l.Error("Error executing command", zap.Error(err))
	}
	return err
}

// Function returns string code of command
func (s SetCurrentProgramScene) Code() string {
	return "SetCurrentProgramScene"
}

// Function returns string description of command
func (s SetCurrentProgramScene) Description() string {
	return "Sets current Program Scene"
}

// Representation of set current preview scene command
type SetCurrentPreviewScene struct {
	PreviewSceneName string `hubman:"preview_scene_name"`
}

// Function provides handler to execute command in OBS
func (s SetCurrentPreviewScene) Run(p ObsProvider, l *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		l.Error("Error executing command", zap.Error(err))
		return err
	}
	_, err = obsClient.Scenes.SetCurrentPreviewScene(&scenes.SetCurrentPreviewSceneParams{
		SceneName: &s.PreviewSceneName,
	})
	if err != nil {
		l.Error("Error executing command", zap.Error(err))
	}
	return err
}

// Function returns string code of command
func (s SetCurrentPreviewScene) Code() string {
	return "SetCurrentPreviewScene"
}

// Function returns string description of command
func (s SetCurrentPreviewScene) Description() string {
	return "Sets current Preview Scene"
}

// Representation of set current program scene by id command
type SetCurrentProgramSceneById struct {
	ProgramSceneId int `hubman:"program_scene_id"`
}

// Function provides handler to execute command in OBS
func (s SetCurrentProgramSceneById) Run(p ObsProvider, l *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		l.Error("Error executing command", zap.Error(err))
		return err
	}
	sceneListResponse, err := obsClient.Scenes.GetSceneList()
	if err != nil {
		l.Error("Error executing command", zap.Error(err))
		return err
	}
	if (s.ProgramSceneId <= 0) || (s.ProgramSceneId > len(sceneListResponse.Scenes)) {
		err = errors.New("Scene id out of range")
		l.Error("Scene id out of range", zap.Int("sceneId", s.ProgramSceneId))
		return err
	}
	programSceneName := sceneListResponse.Scenes[len(sceneListResponse.Scenes)-s.ProgramSceneId].SceneName
	_, err = obsClient.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{
		SceneName: &programSceneName,
	})
	if err != nil {
		l.Error("Error executing command", zap.Error(err))
	}
	return err
}

// Function returns string code of command
func (s SetCurrentProgramSceneById) Code() string {
	return "SetCurrentProgramSceneById"
}

// Function returns string description of command
func (s SetCurrentProgramSceneById) Description() string {
	return "Sets current Program Scene by id"
}

// Representation of set current preview scene by id command
type SetCurrentPreviewSceneById struct {
	PreviewSceneId int `hubman:"preview_scene_id"`
}

// Function provides handler to execute command in OBS
func (s SetCurrentPreviewSceneById) Run(p ObsProvider, l *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		l.Error("Error executing command", zap.Error(err))
		return err
	}
	sceneListResponse, err := obsClient.Scenes.GetSceneList()
	if err != nil {
		l.Error("Error executing command", zap.Error(err))
		return err
	}
	if (s.PreviewSceneId <= 0) || (s.PreviewSceneId > len(sceneListResponse.Scenes)) {
		err = errors.New("Scene id out of range")
		l.Error("Scene id out of range", zap.Int("sceneId", s.PreviewSceneId))
		return err
	}
	previewSceneName := sceneListResponse.Scenes[len(sceneListResponse.Scenes)-s.PreviewSceneId].SceneName
	_, err = obsClient.Scenes.SetCurrentPreviewScene(&scenes.SetCurrentPreviewSceneParams{
		SceneName: &previewSceneName,
	})
	if err != nil {
		l.Error("Error executing command", zap.Error(err))
	}
	return err
}

// Function returns string code of command
func (s SetCurrentPreviewSceneById) Code() string {
	return "SetCurrentPreviewSceneById"
}

// Function returns string description of command
func (s SetCurrentPreviewSceneById) Description() string {
	return "Sets current Preview Scene by id"
}

/*----------------------------- Set SceneItem --------------------------------*/
