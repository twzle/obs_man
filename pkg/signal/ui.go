package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var UISignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[StudioModeStateChanged](),
}

type StudioModeStateChanged struct {
	StudioModeEnabled bool `hubman:"studio_mode_enabled"`
}

func (s StudioModeStateChanged) Code() string {
	return "StudioModeStateChanged"
}

func (s StudioModeStateChanged) Description() string {
	return "Sent when studio mode has been enabled or disabled"
}
