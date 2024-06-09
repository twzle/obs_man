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

// Representation of current program scene changed signal
type CurrentProgramSceneChanged struct {
	SceneName string `hubman:"scene_name"`
}

// Function returns string code of signal
func (c CurrentProgramSceneChanged) Code() string {
	return "CurrentProgramSceneChanged"
}

// Function returns string description of signal
func (c CurrentProgramSceneChanged) Description() string {
	return "Sent when current program scene changes to included scene by name"
}

// Representation of current preview scene changed signal
type CurrentPreviewSceneChanged struct {
	SceneName string `hubman:"scene_name"`
}

// Function returns string code of signal
func (c CurrentPreviewSceneChanged) Code() string {
	return "CurrentPreviewSceneChanged"
}

// Function returns string description of signal
func (c CurrentPreviewSceneChanged) Description() string {
	return "Sent when current preview scene changes to included scene by name"
}

// Representation of current program scene changed by id signal
type CurrentProgramSceneChangedById struct {
	SceneName string `hubman:"scene_name"`
	SceneId   int    `hubman:"scene_id"`
}

// Function returns string code of signal
func (c CurrentProgramSceneChangedById) Code() string {
	return "CurrentProgramSceneChangedById"
}

// Function returns string description of signal
func (c CurrentProgramSceneChangedById) Description() string {
	return "Sent when current program scene changes to included scene by id"
}

// Representation of current preview scene changed by id signal
type CurrentPreviewSceneChangedById struct {
	SceneName string `hubman:"scene_name"`
	SceneId   int    `hubman:"scene_id"`
}

// Function returns string code of signal
func (c CurrentPreviewSceneChangedById) Code() string {
	return "CurrentPreviewSceneChangedById"
}

// Function returns string description of signal
func (c CurrentPreviewSceneChangedById) Description() string {
	return "Sent when current preview scene changes to included scene by id"
}
