package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var StreamSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[StreamStateChanged](),
}

// Representation of stream state changed signal
type StreamStateChanged struct {
	OutputActive bool   `hubman:"output_active"`
	OutputState  string `hubman:"output_state"`
}

// Function returns string code of signal
func (s StreamStateChanged) Code() string {
	return "StreamStateChanged"
}

// Function returns string description of signal
func (s StreamStateChanged) Description() string {
	return "Sent when stream output state changes, for OutputState values see obs docs"
}
