package command

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"github.com/andreykaipov/goobs/api/requests/ui"
	"go.uber.org/zap"
)

// Fucntion providing handlers for command to manage ui with OBS
func ProvideUiCommands(obsProvider ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(SetStudioModeEnabled{}, func(_ core.SerializedCommand, cp ex.CommandParser) error {
			cmd := SetStudioModeEnabled{}
			cp(&cmd)
			return cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
	}
}

// Representation of set studio mode enabled command
type SetStudioModeEnabled struct {
	UseStudioMode bool `hubman:"use_studio_mode"`
}

// Function provides handler to execute command in OBS
func (s SetStudioModeEnabled) Run(p ObsProvider, log *zap.Logger) error {
	obsClient, err := p.Provide()
	if err != nil {
		return logErr(log, "p.Provide", err)
	}
	_, err = obsClient.Ui.SetStudioModeEnabled(&ui.SetStudioModeEnabledParams{
		StudioModeEnabled: &s.UseStudioMode,
	})
	return logErr(log, "obsClient.Ui.SetStudioModeEnabled", err)
}

// Function returns string description of command
func (s SetStudioModeEnabled) Description() string {
	return "Enables or disables studio mode with given property"
}

// Function returns string code of command
func (s SetStudioModeEnabled) Code() string {
	return "SetStudioModeEnabled"
}
