package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var CurrentSceneSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[CurrentProgramSceneChanged](),
	hubman.WithSignal[CurrentPreviewSceneChanged](),
	hubman.WithSignal[CurrentProgramSceneChangedById](),
	hubman.WithSignal[CurrentPreviewSceneChangedById](),
}

type CurrentProgramSceneChanged struct {
	SceneName string `hubman:"scene_name"`
}

func (c CurrentProgramSceneChanged) Code() string {
	return "CurrentProgramSceneChanged"
}

func (c CurrentProgramSceneChanged) Description() string {
	return "Sent when current program scene changes to included scene by name"
}

type CurrentPreviewSceneChanged struct {
	SceneName string `hubman:"scene_name"`
}

func (c CurrentPreviewSceneChanged) Code() string {
	return "CurrentPreviewSceneChanged"
}

func (c CurrentPreviewSceneChanged) Description() string {
	return "Sent when current preview scene changes to included scene by name"
}

type CurrentProgramSceneChangedById struct {
	SceneName string `hubman:"scene_name"`
	SceneId   int    `hubman:"scene_id"`
}

func (c CurrentProgramSceneChangedById) Code() string {
	return "CurrentProgramSceneChangedById"
}

func (c CurrentProgramSceneChangedById) Description() string {
	return "Sent when current program scene changes to included scene by id"
}

type CurrentPreviewSceneChangedById struct {
	SceneName string `hubman:"scene_name"`
	SceneId   int    `hubman:"scene_id"`
}

func (c CurrentPreviewSceneChangedById) Code() string {
	return "CurrentPreviewSceneChangedById"
}

func (c CurrentPreviewSceneChangedById) Description() string {
	return "Sent when current preview scene changes to included scene by id"
}
