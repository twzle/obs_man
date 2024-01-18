package signal

import (
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
)

var SceneItemSignals []func(manipulator.Manipulator) = []func(manipulator.Manipulator){
	hubman.WithSignal[SceneItemEnableStateChanged](),
}

type SceneItemEnableStateChanged struct {
	SceneName        string `hubman:"scene_name"`
	SceneItemId      int    `hubman:"scene_item_id"`
	SceneItemEnabled bool   `hubman:"scene_item_enabled"`
}

func (s SceneItemEnableStateChanged) Code() string {
	return "SceneItemEnableStateChanged"
}

func (s SceneItemEnableStateChanged) Description() string {
	return "Sent when item's enabled state has changed"
}
