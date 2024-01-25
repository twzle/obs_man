package command

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"go.uber.org/zap"
)

func ProvideVirtualCameraCommands(obsManager ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(ToggleVirtualCam{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := ToggleVirtualCam{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(StopVirtualCam{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := StopVirtualCam{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(StartVirtualCam{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := StartVirtualCam{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
	}
}

var _ RunnableCommand = &ToggleVirtualCam{}
var _ RunnableCommand = &StopVirtualCam{}
var _ RunnableCommand = &StartVirtualCam{}

type ToggleVirtualCam struct {
}

func (t ToggleVirtualCam) Code() string {
	return "ToggleVirtualCam"
}

func (t ToggleVirtualCam) Description() string {
	return "Toggles VirtualCam, ex: enabled -> disable, disabled -> enable"
}

func (t ToggleVirtualCam) Run(p ObsProvider, _ *zap.Logger) error {
	obsClient, _, err := p.Provide()
	if err != nil {
		return err
	}
	_, err = obsClient.Outputs.ToggleVirtualCam()
	return err
}

type StopVirtualCam struct {
}

func (s StopVirtualCam) Code() string {
	return "StopVirtualCam"
}

func (s StopVirtualCam) Description() string {
	return "Stops VirtualCam, if it is not running - is no-op"
}

func (s StopVirtualCam) Run(p ObsProvider, _ *zap.Logger) error {
	obsClient, _, err := p.Provide()
	if err != nil {
		return err
	}
	_, err = obsClient.Outputs.StopVirtualCam()
	return err
}

type StartVirtualCam struct {
}

func (s StartVirtualCam) Code() string {
	return "StartVirtualCam"
}

func (s StartVirtualCam) Description() string {
	return "Starts VirtualCam, if it is already running - is no-op"
}

func (s StartVirtualCam) Run(p ObsProvider, _ *zap.Logger) error {
	obsClient, _, err := p.Provide()
	if err != nil {
		return err
	}
	_, err = obsClient.Outputs.StartVirtualCam()
	return err
}
