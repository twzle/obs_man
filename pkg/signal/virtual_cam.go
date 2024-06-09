package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var VirtualCamSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[VirtualCamStateChanged](),
}

// Representation of virtual camera state changed signal
type VirtualCamStateChanged struct {
	OutputActive bool   `hubman:"output_active"`
	OutputState  string `hubman:"output_state"`
}

// Function returns string code of signal
func (v VirtualCamStateChanged) Code() string {
	return "VirtualCamStateChanged"
}

// Function returns string description of signal
func (v VirtualCamStateChanged) Description() string {
	return "Sent when the virtual cam state changes, for OutputState values see obs docs"
}
