package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var InputSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[InputMuteStateChanged](),
	hubman.WithSignal[InputVolumeChanged](),
}

type InputMuteStateChanged struct {
	InputName  string `hubman:"input_name"`
	InputMuted bool   `hubman:"input_muted"`
}

func (i InputMuteStateChanged) Code() string {
	return "InputMuteStateChanged"
}

func (i InputMuteStateChanged) Description() string {
	return "Sent when input with included name mutes or unmutes"
}

type InputVolumeChanged struct {
	InputName      string  `hubman:"input_name"`
	InputVolumeMul float64 `hubman:"input_volume_mul"`
	InputVolumeDb  float64 `hubman:"input_volume_db"`
}

func (i InputVolumeChanged) Code() string {
	return "InputVolumeChanged"
}

func (i InputVolumeChanged) Description() string {
	return "Sent when input with included name changes its volume"
}
