package command

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"github.com/andreykaipov/goobs/api/requests/stream"
	"go.uber.org/zap"
)

func ProvideStreamCommands(obsManager ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(ToggleStream{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := ToggleStream{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(StartStream{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := StartStream{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(StopStream{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := StopStream{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(SendStreamCaption{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := SendStreamCaption{}
			cp(&cmd)
			cmd.Run(obsManager, l.Named(cmd.Code()))
		}),
	}
}

/*----------------------------- Toggle/Start/Stop Stream -------------------*/

var _ RunnableCommand = &ToggleStream{}
var _ RunnableCommand = &StartStream{}
var _ RunnableCommand = &StopStream{}
var _ RunnableCommand = &SendStreamCaption{}

type ToggleStream struct {
}

func (t ToggleStream) Code() string {
	return "ToggleStream"
}

func (t ToggleStream) Description() string {
	return "Toggles Stream, ex: streaming -> stop, stop -> streaming"
}

func (t ToggleStream) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Stream.ToggleStream()
	return logErr(log, "obsClient.Stream.ToggleStream", err)
}

type StartStream struct {
}

func (s StartStream) Code() string {
	return "StartStream"
}

func (s StartStream) Description() string {
	return "Starts Stream, if it is already running is no-op"
}

func (s StartStream) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Stream.StartStream()
	return logErr(log, "obsClient.Stream.StartStream", err)
}

type StopStream struct {
}

func (s StopStream) Code() string {
	return "StopStream"
}

func (s StopStream) Description() string {
	return "Stops Stream, if it is off - is no-op"
}

func (s StopStream) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Stream.StopStream()
	return logErr(log, "obsClient.Stream.StopStream", err)
}

/*----------------------------- SendCaption for stream -----------------------*/

type SendStreamCaption struct {
	StreamCaption string `hubman:"stream_caption"`
}

func (s SendStreamCaption) Code() string {
	return "SendStreamCaption"
}

func (s SendStreamCaption) Description() string {
	return "Sends StreamCaption"
}

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
