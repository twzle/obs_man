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

type SceneTransitionEnded struct {
	TransitionName string `hubman:"transition_name"`
}

func (c SceneTransitionEnded) Code() string {
	return "SceneTransitionEnded"
}

func (c SceneTransitionEnded) Description() string {
	return "Sent when scene transition has completed fully, i.e. not interrupted by user"
}

type SceneTransitionStarted struct {
	TransitionName string `hubman:"transition_name"`
}

func (c SceneTransitionStarted) Code() string {
	return "SceneTransitionStarted"
}

func (c SceneTransitionStarted) Description() string {
	return "Sent when scene transition has started"
}

type SceneTransitionVideoEnded struct {
	TransitionName string `hubman:"transition_name"`
}

func (c SceneTransitionVideoEnded) Code() string {
	return "SceneTransitionStarted"
}

func (c SceneTransitionVideoEnded) Description() string {
	return "Sent when scene transition video complets fully, see obs docs for concrete explanation"
}
