package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var CurrentSceneSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[CurrentProgramSceneChanged](),
	hubman.WithSignal[CurrentPreviewSceneChanged](),
}

type CurrentProgramSceneChanged struct {
	SceneName string `hubman:"scene_name"`
}

func (c CurrentProgramSceneChanged) Code() string {
	return "CurrentProgramSceneChanged"
}

func (c CurrentProgramSceneChanged) Description() string {
	return "Sent when current program scene changes to included scene"
}

type CurrentPreviewSceneChanged struct {
	SceneName string `hubman:"scene_name"`
}

func (c CurrentPreviewSceneChanged) Code() string {
	return "CurrentPreviewSceneChanged"
}

func (c CurrentPreviewSceneChanged) Description() string {
	return "Sent when current preview scene changes to included scene"
}
