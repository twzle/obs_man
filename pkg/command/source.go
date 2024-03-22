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

func ProvideSourceCommands(obsManager ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(SetInputMute{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := SetInputMute{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ToggleInputMute{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := ToggleInputMute{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ToggleSceneItemEnabled{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := ToggleSceneItemEnabled{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ToggleCurrentSceneItemEnabled{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := ToggleCurrentSceneItemEnabled{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
	}
}

var _ RunnableCommand = &SetInputMute{}
var _ RunnableCommand = &ToggleInputMute{}
var _ RunnableCommand = &ToggleSceneItemEnabled{}
var _ RunnableCommand = &ToggleCurrentSceneItemEnabled{}

type SetInputMute struct {
	InputName string `hubman:"input_name"`
	Muted     bool   `hubman:"muted"`
}

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

func (s SetInputMute) Code() string {
	return "SetInputMute"
}

func (s SetInputMute) Description() string {
	return "Sets the audio mute state of an input with given muted property"
}

type ToggleInputMute struct {
	InputName string `hubman:"input_name"`
}

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

func (t ToggleInputMute) Code() string {
	return "ToggleInputMute"
}

func (t ToggleInputMute) Description() string {
	return "Toggles the audio mute state of a given input. Ex true->false, false->true"
}

type ToggleSceneItemEnabled struct {
	SceneItemName string `hubman:"scene_item_name"`
	SceneName     string `hubman:"scene_name"`
}

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

func (t ToggleSceneItemEnabled) Code() string {
	return "ToggleSceneItemEnabled"
}

func (t ToggleSceneItemEnabled) Description() string {
	return "Toggles the scene item enabled state, searches for it using given scene. Ex true->false, false->true"
}

type ToggleCurrentSceneItemEnabled struct {
	SceneItemName string `hubman:"scene_item_name"`
}

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

func (t ToggleCurrentSceneItemEnabled) Code() string {
	return "ToggleCurrentSceneItemEnabled"
}

func (t ToggleCurrentSceneItemEnabled) Description() string {
	return "Toggles the scene item enabled state, searches for it using current scene. Ex true->false, false->true"
}
