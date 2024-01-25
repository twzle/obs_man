package command

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	ex "git.miem.hse.ru/hubman/hubman-lib/executor"
	"github.com/andreykaipov/goobs/api/requests/ui"
	"go.uber.org/zap"
)

func ProvideUiCommands(obsProvider ObsProvider, l *zap.Logger) []func(ex.Executor) {
	return []func(ex.Executor){
		hubman.WithCommand(SetStudioModeEnabled{}, func(_ core.SerializedCommand, cp ex.CommandParser) {
			cmd := SetStudioModeEnabled{}
			cp(&cmd)
			cmd.Run(obsProvider, l.Named(cmd.Code()))
		}),
	}
}

type SetStudioModeEnabled struct {
	UseStudioMode bool `hubman:"use_studio_mode"`
}

func (s SetStudioModeEnabled) Run(p ObsProvider, _ *zap.Logger) error {
	obsClient, _, err := p.Provide()
	if err != nil {
		return err
	}
	_, err = obsClient.Ui.SetStudioModeEnabled(&ui.SetStudioModeEnabledParams{
		StudioModeEnabled: &s.UseStudioMode,
	})
	return err
}

func (s SetStudioModeEnabled) Description() string {
	return "Enables or disables studio mode with given property"
}

func (s SetStudioModeEnabled) Code() string {
	return "SetStudioModeEnabled"
}
