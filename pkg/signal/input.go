package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var InputSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[InputMuteStateChanged](),
	hubman.WithSignal[InputVolumeChanged](),
}

// Representation of input mute state changed signal
type InputMuteStateChanged struct {
	InputName  string `hubman:"input_name"`
	InputMuted bool   `hubman:"input_muted"`
}

// Function returns string code of signal
func (i InputMuteStateChanged) Code() string {
	return "InputMuteStateChanged"
}

// Function returns string description of signal
func (i InputMuteStateChanged) Description() string {
	return "Sent when input with included name mutes or unmutes"
}

// Representation of input volume changed signal
type InputVolumeChanged struct {
	InputName      string  `hubman:"input_name"`
	InputVolumeMul float64 `hubman:"input_volume_mul"`
	InputVolumeDb  float64 `hubman:"input_volume_db"`
}

// Function returns string code of signal
func (i InputVolumeChanged) Code() string {
	return "InputVolumeChanged"
}

// Function returns string description of signal
func (i InputVolumeChanged) Description() string {
	return "Sent when input with included name changes its volume"
}
