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
		hubman.WithCommand(TriggerStudioModeTransition{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := TriggerStudioModeTransition{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(SetCurrentSceneTransition{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := SetCurrentSceneTransition{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(TriggerStudioModeTransitionWithName{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := TriggerStudioModeTransitionWithName{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(SetTBarPosition{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := SetTBarPosition{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
	}
}

var _ RunnableCommand = &SetCurrentSceneTransition{}
var _ RunnableCommand = &TriggerStudioModeTransition{}
var _ RunnableCommand = &TriggerStudioModeTransitionWithName{}
var _ RunnableCommand = &SetTBarPosition{}

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

func (s SetCurrentSceneTransition) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Transitions.SetCurrentSceneTransition(&transitions.SetCurrentSceneTransitionParams{
		TransitionName: &s.TransitionName,
	})
	return logErr(log, "obsClient.Transitions.SetCurrentSceneTransition", err)
}

type TriggerStudioModeTransition struct {
}

func (s TriggerStudioModeTransition) Code() string {
	return "TriggerStudioModeTransition"
}

func (s TriggerStudioModeTransition) Description() string {
	return "Triggers selected in OBS studio mode transition"
}

func (s TriggerStudioModeTransition) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Transitions.TriggerStudioModeTransition()
	return logErr(log, "obsClient.Transitions.TriggerStudioModeTransition:", err)
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

func (s TriggerStudioModeTransitionWithName) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	curTransition, err := obsClient.Transitions.GetCurrentSceneTransition()
	if err != nil {
		return logErr(log, "obsClient.Transitions.GetCurrentSceneTransition", err)
	}
	_, err = obsClient.Transitions.SetCurrentSceneTransition(&transitions.SetCurrentSceneTransitionParams{
		TransitionName: &s.TransitionName,
	})
	if err != nil {
		return logErr(log, "obsClient.Transitions.SetCurrentSceneTransition", err)
	}
	_, err = obsClient.Transitions.TriggerStudioModeTransition()
	if err != nil {
		return logErr(log, "obsClient.Transitions.TriggerStudioModeTransition", err)
	}
	<-time.After(300 * time.Millisecond)
	_, err = obsClient.Transitions.SetCurrentSceneTransition(&transitions.SetCurrentSceneTransitionParams{
		TransitionName: &curTransition.TransitionName,
	})
	return logErr(log, "obsClient.Transitions.SetCurrentSceneTransition", err)
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

type SetTBarPosition struct {
	Position float64 `hubman:"position"` // should be in [0, 1]
	Release  bool    `hubman:"release"`
}

func (s SetTBarPosition) Code() string {
	return "SetTBarPosition"
}

func (s SetTBarPosition) Description() string {
	return "Sets Tbar postion in OBS studio mode from 0 to 1"
}

func (s SetTBarPosition) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}

	_, err = obsClient.Transitions.SetTBarPosition(&transitions.SetTBarPositionParams{
		Position: &s.Position,
		Release:  &s.Release,
	})

	return logErr(log, "obsClient.Transitions.SetTBarPosition", err)
}
