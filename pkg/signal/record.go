package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var RecordSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[RecordStateChanged](),
}

type RecordStateChanged struct {
	OutputActive bool   `hubman:"output_active"`
	OutputState  string `hubman:"output_state"`
	OutputPath   string `hubman:"output_path"`
}

func (r RecordStateChanged) Code() string {
	return "RecordStateChanged"
}

func (r RecordStateChanged) Description() string {
	return "Sent when record output state changes, for OutputState values see obs docs"
}
