package command

import (
	"errors"

	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"github.com/andreykaipov/goobs/api/requests/inputs"
	"github.com/andreykaipov/goobs/api/requests/sceneitems"
	"go.uber.org/zap"
)

// Fucntion providing handlers for command to manage sources with OBS
func ProvideSourceCommands(obsManager ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(SetInputMute{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := SetInputMute{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ToggleInputMute{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := ToggleInputMute{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ToggleSceneItemEnabled{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := ToggleSceneItemEnabled{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ToggleCurrentSceneItemEnabled{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := ToggleCurrentSceneItemEnabled{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
	}
}

var _ RunnableCommand = &SetInputMute{}
var _ RunnableCommand = &ToggleInputMute{}
var _ RunnableCommand = &ToggleSceneItemEnabled{}
var _ RunnableCommand = &ToggleCurrentSceneItemEnabled{}

// Representation of set input mute command
type SetInputMute struct {
	InputName string `hubman:"input_name"`
	Muted     bool   `hubman:"muted"`
}

// Function provides handler to execute command in OBS
func (s SetInputMute) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Inputs.SetInputMute(&inputs.SetInputMuteParams{
		InputName: &s.InputName,
	})
	return logErr(log, "obsClient.Inputs.SetInputMute", err)
}

// Function returns string code of command
func (s SetInputMute) Code() string {
	return "SetInputMute"
}

// Function returns string description of command
func (s SetInputMute) Description() string {
	return "Sets the audio mute state of an input with given muted property"
}

// Representation of toggle input mute command
type ToggleInputMute struct {
	InputName string `hubman:"input_name"`
}

// Function provides handler to execute command in OBS
func (t ToggleInputMute) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Inputs.ToggleInputMute(&inputs.ToggleInputMuteParams{
		InputName: &t.InputName,
	})
	return logErr(log, "obsClient.Inputs.ToggleInputMute", err)
}

// Function returns string code of command
func (t ToggleInputMute) Code() string {
	return "ToggleInputMute"
}

// Function returns string description of command
func (t ToggleInputMute) Description() string {
	return "Toggles the audio mute state of a given input. Ex true->false, false->true"
}


// Representation of toggle scene item enabled command
type ToggleSceneItemEnabled struct {
	SceneItemName string `hubman:"scene_item_name"`
	SceneName     string `hubman:"scene_name"`
}

// Function provides handler to execute command in OBS
func (t ToggleSceneItemEnabled) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	items, err := obsClient.SceneItems.GetSceneItemList(&sceneitems.GetSceneItemListParams{
		SceneName: &t.SceneName,
	})
	if err != nil {
		return err
	}
	for _, item := range items.SceneItems {
		if item.SourceName == t.SceneItemName {
			enabled := !item.SceneItemEnabled
			_, err = obsClient.SceneItems.SetSceneItemEnabled(
				&sceneitems.SetSceneItemEnabledParams{
					SceneName:        &t.SceneName,
					SceneItemId:      &item.SceneItemID,
					SceneItemEnabled: &enabled,
				},
			)
			return logErr(log, "obsClient.SceneItems.SetSceneItemEnabled", err)
		}
	}
	return logErr(log, "cmd.ToggleSceneItemEnabled.Run", errors.New("not found scene item"))
}

// Function returns string code of command
func (t ToggleSceneItemEnabled) Code() string {
	return "ToggleSceneItemEnabled"
}

// Function returns string description of command
func (t ToggleSceneItemEnabled) Description() string {
	return "Toggles the scene item enabled state, searches for it using given scene. Ex true->false, false->true"
}

// Representation of toggle current scene item enabled command
type ToggleCurrentSceneItemEnabled struct {
	SceneItemName string `hubman:"scene_item_name"`
}

// Function provides handler to execute command in OBS
func (t ToggleCurrentSceneItemEnabled) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	curScene, err := obsClient.Scenes.GetCurrentProgramScene()
	if err != nil {
		return logErr(log, "obsClient.Scenes.GetCurrentProgramScene", err)
	}
	items, err := obsClient.SceneItems.GetSceneItemList(&sceneitems.GetSceneItemListParams{
		SceneName: &curScene.CurrentProgramSceneName,
	})
	if err != nil {
		return logErr(log, "obsClient.SceneItems.GetSceneItemList", err)
	}
	for _, item := range items.SceneItems {
		if item.SourceName == t.SceneItemName {
			enabled := !item.SceneItemEnabled
			_, err = obsClient.SceneItems.SetSceneItemEnabled(
				&sceneitems.SetSceneItemEnabledParams{
					SceneName:        &curScene.CurrentProgramSceneName,
					SceneItemId:      &item.SceneItemID,
					SceneItemEnabled: &enabled,
				},
			)
			return logErr(log, "obsClient.SceneItems.SetSceneItemEnabled", err)
		}
	}
	return logErr(log, "cmd.ToggleCurrentSceneItemEnabled.Run", errors.New("not found scene item"))
}

// Function returns string code of command
func (t ToggleCurrentSceneItemEnabled) Code() string {
	return "ToggleCurrentSceneItemEnabled"
}

// Function returns string description of command
func (t ToggleCurrentSceneItemEnabled) Description() string {
	return "Toggles the scene item enabled state, searches for it using current scene. Ex true->false, false->true"
}
