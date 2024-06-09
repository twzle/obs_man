package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var RecordSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[RecordStateChanged](),
}

// Representation of record state changed signal
type RecordStateChanged struct {
	OutputActive bool   `hubman:"output_active"`
	OutputState  string `hubman:"output_state"`
	OutputPath   string `hubman:"output_path"`
}

// Function returns string code of signal
func (r RecordStateChanged) Code() string {
	return "RecordStateChanged"
}

// Function returns string description of signal
func (r RecordStateChanged) Description() string {
	return "Sent when record output state changes, for OutputState values see obs docs"
}
