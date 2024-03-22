package command

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"go.uber.org/zap"
)

func ProvideRecordCommands(obsProvider ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(StartRecord{}, func(command core.SerializedCommand, cp ex.CommandParser) {
			cmd := StartRecord{}
			cp(&cmd)
			cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(StopRecord{}, func(command core.SerializedCommand, cp ex.CommandParser) {
			cmd := StopRecord{}
			cp(&cmd)
			cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ToggleRecord{}, func(command core.SerializedCommand, cp ex.CommandParser) {
			cmd := ToggleRecord{}
			cp(&cmd)
			cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),

		hubman.WithCommand(PauseRecord{}, func(command core.SerializedCommand, cp ex.CommandParser) {
			cmd := PauseRecord{}
			cp(&cmd)
			cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ResumeRecord{}, func(command core.SerializedCommand, cp ex.CommandParser) {
			cmd := ResumeRecord{}
			cp(&cmd)
			cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ToggleRecordPause{}, func(command core.SerializedCommand, cp ex.CommandParser) {
			cmd := ToggleRecordPause{}
			cp(&cmd)
			cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
	}
}

/*----------------------------- Start/Stop/Toggle Record ---------------------*/

var _ RunnableCommand = &StartRecord{}
var _ RunnableCommand = &StopRecord{}
var _ RunnableCommand = &ToggleRecord{}

type StartRecord struct {
}

func (s StartRecord) Code() string {
	return "StartRecord"
}

func (s StartRecord) Description() string {
	return "Starts record, if it is already started - is noop. Similar to start record button"
}

func (s StartRecord) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Record.StartRecord()
	return logErr(log, "obsClient.Record.StartRecord", err)
}

type StopRecord struct {
}

func (s StopRecord) Code() string {
	return "StopRecord"
}

func (s StopRecord) Description() string {
	return "Toggles off record, if it is off - is noop. Similar to stop record button"
}

func (s StopRecord) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Record.StopRecord()
	return logErr(log, "obsClient.Record.StopRecord", err)
}

type ToggleRecord struct {
}

func (t ToggleRecord) Code() string {
	return "ToggleRecord"
}

func (t ToggleRecord) Description() string {
	return "Toggles Record, ex recording -> stop off, no recording -> start it"
}

func (t ToggleRecord) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Record.ToggleRecord()
	return logErr(log, "obsClient.Record.ToggleRecord", err)
}

/*----------------------------- Pause/Resume/Toggle PauseRecord --------------*/

var _ RunnableCommand = &PauseRecord{}
var _ RunnableCommand = &ResumeRecord{}
var _ RunnableCommand = &ToggleRecordPause{}

type PauseRecord struct {
}

func (p PauseRecord) Code() string {
	return "PauseRecord"
}

func (p PauseRecord) Description() string {
	return "Pauses current recording, no-op if obs is not recording now"
}

func (p PauseRecord) Run(pr ObsProvider, log *zap.Logger) error {
	obsClient, err := pr.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Record.PauseRecord()
	return logErr(log, "obsClient.Record.PauseRecord", err)
}

type ResumeRecord struct {
}

func (r ResumeRecord) Code() string {
	return "ResumeRecord"
}

func (r ResumeRecord) Description() string {
	return "Resumes Record"
}

func (r ResumeRecord) Run(pr ObsProvider, log *zap.Logger) error {
	obsClient, err := pr.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Record.ResumeRecord()
	return logErr(log, "obsClient.Record.ResumeRecord", err)
}

type ToggleRecordPause struct {
}

func (t ToggleRecordPause) Code() string {
	return "ToggleRecordPause"
}

func (t ToggleRecordPause) Description() string {
	return "Toggles RecordPause"
}

func (t ToggleRecordPause) Run(pr ObsProvider, log *zap.Logger) error {
	obsClient, err := pr.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Record.ToggleRecordPause()
	return logErr(log, "obsClient.Record.ToggleRecordPause", err)
}
