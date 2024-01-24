package command

import (
	"time"

	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"github.com/andreykaipov/goobs/api/requests/transitions"
	"go.uber.org/zap"
)

func ProvideTransitionCommands(obsManager ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(TriggerStudioModeTransition{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := TriggerStudioModeTransition{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(SetCurrentSceneTransition{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := SetCurrentSceneTransition{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(TriggerStudioModeTransitionWithName{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := TriggerStudioModeTransitionWithName{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
	}
}

var _ RunnableCommand = &SetCurrentSceneTransition{}
var _ RunnableCommand = &TriggerStudioModeTransition{}
var _ RunnableCommand = &TriggerStudioModeTransitionWithName{}

// TODO add commands to executor
type SetSceneSceneTransitionOverride struct {
	SceneName          string  `hubman:"scene_name"`
	TransitionDuration float64 `hubman:"transition_duration"`
	TransitionName     string  `hubman:"transition_name"`
}

func (s SetSceneSceneTransitionOverride) Code() string {
	return "SetSceneSceneTransitionOverride"
}

func (s SetSceneSceneTransitionOverride) Description() string {
	return "Sets SceneScene Transition Override"
}

type SetCurrentSceneTransition struct {
	TransitionName string `hubman:"transition_name"`
}

func (s SetCurrentSceneTransition) Code() string {
	return "SetCurrentSceneTransition"
}

func (s SetCurrentSceneTransition) Description() string {
	return "Sets Current Scene Transition"
}

func (s SetCurrentSceneTransition) Run(p ObsProvider, _ *zap.Logger) error {
	obsClient, _, err := p.Provide()
	if err != nil {
		return err
	}
	_, err = obsClient.Transitions.SetCurrentSceneTransition(&transitions.SetCurrentSceneTransitionParams{
		TransitionName: s.TransitionName,
	})
	return err
}

type TriggerStudioModeTransition struct {
}

func (s TriggerStudioModeTransition) Code() string {
	return "TriggerStudioModeTransition"
}

func (s TriggerStudioModeTransition) Description() string {
	return "Triggers selected in OBS studio mode transition"
}

func (s TriggerStudioModeTransition) Run(p ObsProvider, _ *zap.Logger) error {
	obsClient, _, err := p.Provide()
	if err != nil {
		return err
	}
	_, err = obsClient.Transitions.TriggerStudioModeTransition()
	return err
}

type TriggerStudioModeTransitionWithName struct {
	TransitionName string `hubman:"transition_name"`
}

func (s TriggerStudioModeTransitionWithName) Code() string {
	return "TriggerStudioModeTransitionWithName"
}

func (s TriggerStudioModeTransitionWithName) Description() string {
	return "Triggers studio mode transition with name included in command"
}

func (s TriggerStudioModeTransitionWithName) Run(p ObsProvider, l *zap.Logger) error {
	obsClient, _, err := p.Provide()
	if err != nil {
		return err
	}
	curTransition, err := obsClient.Transitions.GetCurrentSceneTransition()
	if err != nil {
		return err
	}
	_, err = obsClient.Transitions.SetCurrentSceneTransition(&transitions.SetCurrentSceneTransitionParams{
		TransitionName: s.TransitionName,
	})
	if err != nil {
		return err
	}
	_, err = obsClient.Transitions.TriggerStudioModeTransition()
	if err != nil {
		return err
	}
	<-time.After(300 * time.Millisecond)
	_, err = obsClient.Transitions.SetCurrentSceneTransition(&transitions.SetCurrentSceneTransitionParams{
		TransitionName: curTransition.TransitionName,
	})
	return err
}

type SetCurrentSceneCollection struct {
	SceneCollectionName string `hubman:"scene_collection_name"`
}

func (s SetCurrentSceneCollection) Code() string {
	return "SetCurrentSceneCollection"
}

func (s SetCurrentSceneCollection) Description() string {
	return "Sets Current Scene Collection"
}

type SetSceneItemBlendMode struct {
	SceneItemBlendMode string  `hubman:"scene_name"`
	SceneItemId        float64 `hubman:"scene_item_id"`
	SceneName          string  `hubman:"scene_name"`
}

func (s SetSceneItemBlendMode) Code() string {
	return "SetSceneItemBlendMode"
}

func (s SetSceneItemBlendMode) Description() string {
	return "Sets Scene ItemBlendMode"
}
