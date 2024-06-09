package command

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"github.com/andreykaipov/goobs/api/requests/stream"
	"go.uber.org/zap"
)

// Fucntion providing handlers for command to manage stream with OBS
func ProvideStreamCommands(obsManager ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(ToggleStream{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := ToggleStream{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(StartStream{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := StartStream{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(StopStream{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := StopStream{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(SendStreamCaption{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := SendStreamCaption{}
			cp(&cmd)
			return cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
	}
}

/*----------------------------- Toggle/Start/Stop Stream -------------------*/

var _ RunnableCommand = &ToggleStream{}
var _ RunnableCommand = &StartStream{}
var _ RunnableCommand = &StopStream{}
var _ RunnableCommand = &SendStreamCaption{}

// Representation of toggle stream command
type ToggleStream struct {
}

// Function returns string code of command
func (t ToggleStream) Code() string {
	return "ToggleStream"
}

// Function returns string description of command
func (t ToggleStream) Description() string {
	return "Toggles Stream, ex: streaming -> stop, stop -> streaming"
}

// Function provides handler to execute command in OBS
func (t ToggleStream) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Stream.ToggleStream()
	return logErr(log, "obsClient.Stream.ToggleStream", err)
}

// Representation of start stream command
type StartStream struct {
}

// Function returns string code of command
func (s StartStream) Code() string {
	return "StartStream"
}

// Function returns string description of command
func (s StartStream) Description() string {
	return "Starts Stream, if it is already running is no-op"
}

// Function provides handler to execute command in OBS
func (s StartStream) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Stream.StartStream()
	return logErr(log, "obsClient.Stream.StartStream", err)
}

// Representation of stop stream command
type StopStream struct {
}

// Function returns string code of command
func (s StopStream) Code() string {
	return "StopStream"
}

// Function returns string description of command
func (s StopStream) Description() string {
	return "Stops Stream, if it is off - is no-op"
}

// Function provides handler to execute command in OBS
func (s StopStream) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Stream.StopStream()
	return logErr(log, "obsClient.Stream.StopStream", err)
}

/*----------------------------- SendCaption for stream -----------------------*/

// Representation of send stream caption command
type SendStreamCaption struct {
	StreamCaption string `hubman:"stream_caption"`
}

// Function returns string code of command
func (s SendStreamCaption) Code() string {
	return "SendStreamCaption"
}

// Function returns string description of command
func (s SendStreamCaption) Description() string {
	return "Sends StreamCaption"
}

// Function provides handler to execute command in OBS
func (s SendStreamCaption) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Stream.SendStreamCaption(&stream.SendStreamCaptionParams{
		CaptionText: &s.StreamCaption,
	})
	return logErr(log, "obsClient.Stream.SendStreamCaption", err)
}
