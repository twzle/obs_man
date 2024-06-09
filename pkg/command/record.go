package command

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"go.uber.org/zap"
)

// Fucntion providing handlers for command to manage record process with OBS
func ProvideRecordCommands(obsProvider ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(StartRecord{}, func(command core.SerializedCommand, cp ex.CommandParser) error {
			cmd := StartRecord{}
			cp(&cmd)
			return cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(StopRecord{}, func(command core.SerializedCommand, cp ex.CommandParser) error {
			cmd := StopRecord{}
			cp(&cmd)
			return cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ToggleRecord{}, func(command core.SerializedCommand, cp ex.CommandParser) error {
			cmd := ToggleRecord{}
			cp(&cmd)
			return cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),

		hubman.WithCommand(PauseRecord{}, func(command core.SerializedCommand, cp ex.CommandParser) error {
			cmd := PauseRecord{}
			cp(&cmd)
			return cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ResumeRecord{}, func(command core.SerializedCommand, cp ex.CommandParser) error {
			cmd := ResumeRecord{}
			cp(&cmd)
			return cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
		hubman.WithCommand(ToggleRecordPause{}, func(command core.SerializedCommand, cp ex.CommandParser) error {
			cmd := ToggleRecordPause{}
			cp(&cmd)
			return cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
	}
}

/*----------------------------- Start/Stop/Toggle Record ---------------------*/

var _ RunnableCommand = &StartRecord{}
var _ RunnableCommand = &StopRecord{}
var _ RunnableCommand = &ToggleRecord{}

// Representation of start record command
type StartRecord struct {
}

// Function returns string code of command
func (s StartRecord) Code() string {
	return "StartRecord"
}

// Function returns string description of command
func (s StartRecord) Description() string {
	return "Starts record, if it is already started - is noop. Similar to start record button"
}

// Function provides handler to execute command in OBS
func (s StartRecord) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Record.StartRecord()
	return logErr(log, "obsClient.Record.StartRecord", err)
}

// Representation of stop record command
type StopRecord struct {
}

// Function returns string code of command
func (s StopRecord) Code() string {
	return "StopRecord"
}

// Function returns string description of command
func (s StopRecord) Description() string {
	return "Toggles off record, if it is off - is noop. Similar to stop record button"
}

// Function provides handler to execute command in OBS
func (s StopRecord) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Record.StopRecord()
	return logErr(log, "obsClient.Record.StopRecord", err)
}

// Representation of toggle record command
type ToggleRecord struct {
}

// Function returns string code of command
func (t ToggleRecord) Code() string {
	return "ToggleRecord"
}

// Function returns string description of command
func (t ToggleRecord) Description() string {
	return "Toggles Record, ex recording -> stop off, no recording -> start it"
}

// Function provides handler to execute command in OBS
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

// Representation of pause record command
type PauseRecord struct {
}

// Function returns string code of command
func (p PauseRecord) Code() string {
	return "PauseRecord"
}

// Function returns string description of command
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

// Representation of resume record command
type ResumeRecord struct {
}

// Function returns string code of command
func (r ResumeRecord) Code() string {
	return "ResumeRecord"
}

// Function returns string description of command
func (r ResumeRecord) Description() string {
	return "Resumes Record"
}

// Function provides handler to execute command in OBS
func (r ResumeRecord) Run(pr ObsProvider, log *zap.Logger) error {
	obsClient, err := pr.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Record.ResumeRecord()
	return logErr(log, "obsClient.Record.ResumeRecord", err)
}

// Representation of toggle record pause command
type ToggleRecordPause struct {
}

// Function returns string code of command
func (t ToggleRecordPause) Code() string {
	return "ToggleRecordPause"
}

// Function returns string description of command
func (t ToggleRecordPause) Description() string {
	return "Toggles RecordPause"
}

// Function provides handler to execute command in OBS
func (t ToggleRecordPause) Run(pr ObsProvider, log *zap.Logger) error {
	obsClient, err := pr.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Record.ToggleRecordPause()
	return logErr(log, "obsClient.Record.ToggleRecordPause", err)
}
