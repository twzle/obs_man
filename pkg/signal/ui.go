package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var UISignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[StudioModeStateChanged](),
}

// Representation of studio mode state changed signal
type StudioModeStateChanged struct {
	StudioModeEnabled bool `hubman:"studio_mode_enabled"`
}

// Function returns string code of signal
func (s StudioModeStateChanged) Code() string {
	return "StudioModeStateChanged"
}

// Function returns string description of signal
func (s StudioModeStateChanged) Description() string {
	return "Sent when studio mode has been enabled or disabled"
}
