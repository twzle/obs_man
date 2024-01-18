package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var ReplayBufferSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[ReplayBufferSaved](),
	hubman.WithSignal[ReplayBufferStateChanged](),
}

type ReplayBufferStateChanged struct {
	OutputActive bool   `hubman:"output_active"`
	OutputState  string `hubman:"output_state"`
}

func (r ReplayBufferStateChanged) Code() string {
	return "ReplayBufferStateChanged"
}

func (r ReplayBufferStateChanged) Description() string {
	return "Sent when the replay buffer state changes, for OutputState values see obs docs"
}

type ReplayBufferSaved struct {
	SavedReplayPath string `hubman:"saved_replay_path"`
}

func (r ReplayBufferSaved) Code() string {
	return "ReplayBufferSaved"
}

func (r ReplayBufferSaved) Description() string {
	return "Sent when the replay buffer has been saved to included path"
}
