package obsman

import (
	"errors"
)

type AgentConf struct {
	ExecutorPort uint16 `yaml:"server_port"`
	RedisUrl     string `yaml:"redis_url"`
}

type ObsConf struct {
	HostPort string `yaml:"host_port"`
	Password string `yaml:"password"`
}

type Conf struct {
	Agent   AgentConf `yaml:"agent"`
	ObsConf ObsConf   `yaml:"obs"`
}

func (o ObsConf) Validate() error {
	if o.HostPort == "" {
		return errors.New("empty host port")
	}
	return nil
}

type SetScene struct {
	SceneName string `hubman:"scene_name"`
}

func (s SetScene) Code() string {
	return "SetScene"
}

func (s SetScene) Description() string {
	return "Sets scene with given name active in obs"
}

type StartRecord struct {
}

func (s StartRecord) Code() string {
	return "StartRecord"
}

func (s StartRecord) Description() string {
	return "Toggles record, if it is started - is noop. Similar to start record button"
}

type StopRecord struct {
}

func (s StopRecord) Code() string {
	return "StopRecord"
}

func (s StopRecord) Description() string {
	return "Toggles off record, if it is off - is noop. Similar to stop record button"
}

//-----------------------------------------------

type SetCurrentProgramScene struct {
	ProgramSceneName string `hubman:"program_scene_name"`
}

func (s SetCurrentProgramScene) Code() string {
	return "SetCurrentProgramScene"
}

func (s SetCurrentProgramScene) Description() string {
	return "Sets current Program Scene"
}

type SetCurrentPreviewScene struct {
	PreviewSceneName string `hubman:"preview_scene_name"`
}

func (s SetCurrentPreviewScene) Code() string {
	return "SetCurrentPreviewScene"
}

func (s SetCurrentPreviewScene) Description() string {
	return "Sets current Preview Scene"
}

type CreateScene struct {
	SceneName string `hubman:"scene_name"`
}

func (s CreateScene) Code() string {
	return "CreateScene"
}

func (s CreateScene) Description() string {
	return "Creates scene"
}

type SetSceneName struct {
	SceneName string `hubman:"scene_name"`
}

func (s SetSceneName) Code() string {
	return "SetSceneName"
}

func (s SetSceneName) Description() string {
	return "Sets scene name"
}

type SetSceneSceneTransitionOverride struct {
	SceneName          string  `hubman:"scene_name"`
	TransitionDuration float64 `hubman:"transition_duration"`
	TransitionName     string  `hubman:"transition_name"`
}

func (s SetSceneSceneTransitionOverride) Code() string {
	return "SetSceneSceneTransitionOverride"
}

func (s SetSceneSceneTransitionOverride) Description() string {
	return "Sets SceneScene Transition Override"
}

type SetCurrentSceneTransition struct {
	TransitionName string `hubman:"transition_name"`
}

func (s SetCurrentSceneTransition) Code() string {
	return "SetCurrentSceneTransition"
}

func (s SetCurrentSceneTransition) Description() string {
	return "Sets Current Scene Transition"
}

type TriggerStudioModeTransition struct {
}

func (s TriggerStudioModeTransition) Code() string {
	return "TriggerStudioModeTransition"
}

func (s TriggerStudioModeTransition) Description() string {
	return "Triggers Studio Mode Transition"
}

type SetCurrentSceneCollection struct {
	SceneCollectionName string `hubman:"scene_collection_name"`
}

func (s SetCurrentSceneCollection) Code() string {
	return "SetCurrentSceneCollection"
}

func (s SetCurrentSceneCollection) Description() string {
	return "Sets Current Scene Collection"
}

type SetSceneItemBlendMode struct {
	SceneItemBlendMode string  `hubman:"scene_name"`
	SceneItemId        float64 `hubman:"scene_item_id"`
	SceneName          string  `hubman:"scene_name"`
}

func (s SetSceneItemBlendMode) Code() string {
	return "SetSceneItemBlendMode"
}

func (s SetSceneItemBlendMode) Description() string {
	return "Sets Scene ItemBlendMode"
}

type SetSceneItemTransform struct {
	Alignment       float64 `hubman:"alignment"`
	BoundsAlignment float64 `hubman:"boundsAlignment"`
	BoundsHeight    float64 `hubman:"boundsHeight"`
	BoundsType      string  `hubman:"boundsType"`
	BoundsWidth     float64 `hubman:"boundsWidth"`
	CropBottom      float64 `hubman:"cropBottom"`
	CropLeft        float64 `hubman:"cropLeft"`
	CropRight       float64 `hubman:"cropRight"`
	CropTop         float64 `hubman:"cropTop"`
	Height          float64 `hubman:"height"`
	PositionX       float64 `hubman:"positionX"`
	PositionY       float64 `hubman:"positionY"`
	Rotation        float64 `hubman:"rotation"`
	ScaleX          float64 `hubman:"scaleX"`
	ScaleY          float64 `hubman:"scaleY"`
	SourceHeight    float64 `hubman:"sourceHeight"`
	SourceWidth     float64 `hubman:"sourceWidth"`
	Width           float64 `hubman:"width"`
	SceneItemId     float64 `hubman:"scene_item_id"`
	SceneName       string  `hubman:"scene_name"`
}

func (s SetSceneItemTransform) Code() string {
	return "SetSceneItemTransform"
}

func (s SetSceneItemTransform) Description() string {
	return "Sets Scene ItemTransform"
}

type RemoveSceneItem struct {
	SceneItemId float64 `hubman:"scene_item_id"`
	SceneName   string  `hubman:"scene_name"`
}

func (s RemoveSceneItem) Code() string {
	return "RemoveSceneItem"
}

func (s RemoveSceneItem) Description() string {
	return "Removes SceneItem"
}

type ToggleVirtualCam struct {
}

func (s ToggleVirtualCam) Code() string {
	return "ToggleVirtualCam"
}

func (s ToggleVirtualCam) Description() string {
	return "Toggles VirtualCam"
}

type StopVirtualCam struct {
}

func (s StopVirtualCam) Code() string {
	return "StopVirtualCam"
}

func (s StopVirtualCam) Description() string {
	return "Stops VirtualCam"
}

type StartVirtualCam struct {
}

func (s StartVirtualCam) Code() string {
	return "StartVirtualCam"
}

func (s StartVirtualCam) Description() string {
	return "Starts VirtualCam"
}

type ToggleStream struct {
}

func (s ToggleStream) Code() string {
	return "ToggleStream"
}

func (s ToggleStream) Description() string {
	return "Toggles Stream"
}

type StartStream struct {
}

func (s StartStream) Code() string {
	return "StartStream"
}

func (s StartStream) Description() string {
	return "Starts Stream"
}

type StopStream struct {
}

func (s StopStream) Code() string {
	return "StopStream"
}

func (s StopStream) Description() string {
	return "Stops Stream"
}

type SendStreamCaption struct {
	StreamCaption string `hubman:"stream_caption"`
}

func (s SendStreamCaption) Code() string {
	return "SendStreamCaption"
}

func (s SendStreamCaption) Description() string {
	return "Sends StreamCaption"
}

type PauseRecord struct {
}

func (s PauseRecord) Code() string {
	return "PauseRecord"
}

func (s PauseRecord) Description() string {
	return "Pauses Record"
}

type ResumeRecord struct {
}

func (s ResumeRecord) Code() string {
	return "ResumeRecord"
}

func (s ResumeRecord) Description() string {
	return "Resumes Record"
}

type ToggleRecord struct {
}

func (s ToggleRecord) Code() string {
	return "ToggleRecord"
}

func (s ToggleRecord) Description() string {
	return "Toggles Record"
}

type ToggleRecordPause struct {
}

func (s ToggleRecordPause) Code() string {
	return "ToggleRecordPause"
}

func (s ToggleRecordPause) Description() string {
	return "Toggles RecordPause"
}

/* OBS GET commands

type GetSceneList struct {
}

func (s GetSceneList) Code() string {
	return "GetSceneList"
}

func (s GetSceneList) Description() string {
	return "Gets scene list in obs"
}

type GetGroupList struct {
}

func (s GetGroupList) Code() string {
	return "GetGroupList"
}

func (s GetGroupList) Description() string {
	return "GetGroupList"
}

type GetCurrentProgramScene struct {
}

func (s GetCurrentProgramScene) Code() string {
	return "GetCurrentProgramScene"
}

func (s GetCurrentProgramScene) Description() string {
	return "GetCurrentProgramScene"
}

type GetCurrentPreviewScene struct {
}

func (s GetCurrentPreviewScene) Code() string {
	return "GetCurrentPreviewScene"
}

func (s GetCurrentPreviewScene) Description() string {
	return "GetCurrentPreviewScene"
}

type GetSceneSceneTransitionOverride struct {
}

func (s GetSceneSceneTransitionOverride) Code() string {
	return "GetSceneSceneTransitionOverride"
}

func (s GetSceneSceneTransitionOverride) Description() string {
	return "GetSceneSceneTransitionOverride"
}

type GetTransitionKindList struct {
}

func (s GetTransitionKindList) Code() string {
	return "GetTransitionKindList"
}

func (s GetTransitionKindList) Description() string {
	return "GetTransitionKindList"
}

type GetCurrentSceneTransition struct {
}

func (s GetCurrentSceneTransition) Code() string {
	return "GetCurrentSceneTransition"
}

func (s GetCurrentSceneTransition) Description() string {
	return "GetCurrentSceneTransition"
}

type GetSceneCollectionList struct {
}

func (s GetSceneCollectionList) Code() string {
	return "GetSceneCollectionList"
}

func (s GetSceneCollectionList) Description() string {
	return "GetSceneCollectionList"
}

type GetSceneItemBlendMode struct {
}

func (s GetSceneItemBlendMode) Code() string {
	return "GetSceneItemBlendMode"
}

func (s GetSceneItemBlendMode) Description() string {
	return "GetSceneItemBlendMode"
}

type GetSceneItemEnabled struct {
}

func (s GetSceneItemEnabled) Code() string {
	return "GetSceneItemEnabled"
}

func (s GetSceneItemEnabled) Description() string {
	return "GetSceneItemEnabled"
}

type GetSceneItemTransform struct {
}

func (s GetSceneItemTransform) Code() string {
	return "GetSceneItemTransform"
}

func (s GetSceneItemTransform) Description() string {
	return "GetSceneItemTransform"
}

type GetVirtualCamStatus struct {
}

func (s GetVirtualCamStatus) Code() string {
	return "GetVirtualCamStatus"
}

func (s GetVirtualCamStatus) Description() string {
	return "GetVirtualCamStatus"
}

type GetStreamStatus struct {
}

func (s GetStreamStatus) Code() string {
	return "GetStreamStatus"
}

func (s GetStreamStatus) Description() string {
	return "GetStreamStatus"
}

type GetRecordStatus struct {
}

func (s GetRecordStatus) Code() string {
	return "GetRecordStatus"
}

func (s GetRecordStatus) Description() string {
	return "GetRecordStatus"
}
*/
//------------------------------------------------
