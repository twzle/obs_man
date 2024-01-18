package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var VirtualCamSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[VirtualCamStateChanged](),
}

type VirtualCamStateChanged struct {
	OutputActive bool   `hubman:"output_active"`
	OutputState  string `hubman:"output_state"`
}

func (v VirtualCamStateChanged) Code() string {
	return "VirtualCamStateChanged"
}

func (v VirtualCamStateChanged) Description() string {
	return "Sent when the virtual cam state changes, for OutputState values see obs docs"
}
