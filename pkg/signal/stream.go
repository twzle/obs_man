package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var StreamSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[StreamStateChanged](),
}

type StreamStateChanged struct {
	OutputActive bool   `hubman:"output_active"`
	OutputState  string `hubman:"output_state"`
}

func (s StreamStateChanged) Code() string {
	return "StreamStateChanged"
}

func (s StreamStateChanged) Description() string {
	return "Sent when stream output state changes, for OutputState values see obs docs"
}
