package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var TransitionSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[SceneTransitionEnded](),
	hubman.WithSignal[SceneTransitionStarted](),
	hubman.WithSignal[SceneTransitionVideoEnded](),
}

// Representation of scene transtion ended signal
type SceneTransitionEnded struct {
	TransitionName string `hubman:"transition_name"`
}

// Function returns string code of signal
func (c SceneTransitionEnded) Code() string {
	return "SceneTransitionEnded"
}

// Function returns string description of signal
func (c SceneTransitionEnded) Description() string {
	return "Sent when scene transition has completed fully, i.e. not interrupted by user"
}

// Representation of scene transtion started signal
type SceneTransitionStarted struct {
	TransitionName string `hubman:"transition_name"`
}

// Function returns string code of signal
func (c SceneTransitionStarted) Code() string {
	return "SceneTransitionStarted"
}

// Function returns string description of signal
func (c SceneTransitionStarted) Description() string {
	return "Sent when scene transition has started"
}

// Representation of scene transtion video ended signal
type SceneTransitionVideoEnded struct {
	TransitionName string `hubman:"transition_name"`
}

// Function returns string code of signal
func (c SceneTransitionVideoEnded) Code() string {
	return "SceneTransitionStarted"
}

// Function returns string description of signal
func (c SceneTransitionVideoEnded) Description() string {
	return "Sent when scene transition video complets fully, see obs docs for concrete explanation"
}
