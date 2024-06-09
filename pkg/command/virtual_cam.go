package command

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"go.uber.org/zap"
)

// Fucntion providing handlers for command to manage virtual camera with OBS
func ProvideVirtualCameraCommands(obsManager ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(ToggleVirtualCam{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := ToggleVirtualCam{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(StopVirtualCam{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := StopVirtualCam{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(StartVirtualCam{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := StartVirtualCam{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
	}
}

var _ RunnableCommand = &ToggleVirtualCam{}
var _ RunnableCommand = &StopVirtualCam{}
var _ RunnableCommand = &StartVirtualCam{}

// Representation of toggle virtual camera command
type ToggleVirtualCam struct {
}

// Function returns string code of command
func (t ToggleVirtualCam) Code() string {
	return "ToggleVirtualCam"
}

// Function returns string description of command
func (t ToggleVirtualCam) Description() string {
	return "Toggles VirtualCam, ex: enabled -> disable, disabled -> enable"
}

// Function provides handler to execute command in OBS
func (t ToggleVirtualCam) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Outputs.ToggleVirtualCam()
	return logErr(log, "obsClient.Outputs.ToggleVirtualCam", err)
}

// Representation of stop virtual camera command
type StopVirtualCam struct {
}

// Function returns string code of command
func (s StopVirtualCam) Code() string {
	return "StopVirtualCam"
}

// Function returns string description of command
func (s StopVirtualCam) Description() string {
	return "Stops VirtualCam, if it is not running - is no-op"
}

// Function provides handler to execute command in OBS
func (s StopVirtualCam) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Outputs.StopVirtualCam()
	return logErr(log, "obsClient.Outputs.StopVirtualCam", err)
}

// Representation of start virtual camera command
type StartVirtualCam struct {
}

// Function returns string code of command
func (s StartVirtualCam) Code() string {
	return "StartVirtualCam"
}

// Function returns string description of command
func (s StartVirtualCam) Description() string {
	return "Starts VirtualCam, if it is already running - is no-op"
}

// Function provides handler to execute command in OBS
func (s StartVirtualCam) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Outputs.StartVirtualCam()
	return logErr(log, "obsClient.Outputs.StartVirtualCam", err)
}
