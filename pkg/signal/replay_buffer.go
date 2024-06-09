package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var ReplayBufferSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[ReplayBufferSaved](),
	hubman.WithSignal[ReplayBufferStateChanged](),
}

// Representation of replay buffer state changed signal
type ReplayBufferStateChanged struct {
	OutputActive bool   `hubman:"output_active"`
	OutputState  string `hubman:"output_state"`
}

// Function returns string code of signal
func (r ReplayBufferStateChanged) Code() string {
	return "ReplayBufferStateChanged"
}

// Function returns string description of signal
func (r ReplayBufferStateChanged) Description() string {
	return "Sent when the replay buffer state changes, for OutputState values see obs docs"
}

// Representation of replay buffer saved signal
type ReplayBufferSaved struct {
	SavedReplayPath string `hubman:"saved_replay_path"`
}

// Function returns string code of signal
func (r ReplayBufferSaved) Code() string {
	return "ReplayBufferSaved"
}

// Function returns string description of signal
func (r ReplayBufferSaved) Description() string {
	return "Sent when the replay buffer has been saved to included path"
}
