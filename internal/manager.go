package obsman

import (
	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/config"
	"github.com/andreykaipov/goobs/api/requests/sceneitems"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/andreykaipov/goobs/api/requests/stream"
	"github.com/andreykaipov/goobs/api/requests/transitions"
	"github.com/andreykaipov/goobs/api/typedefs"
	"go.uber.org/zap"
	"io"
)

type Manager interface {
	io.Closer
	DoSetScene(SetScene)
	DoStartRecord(StartRecord)
	DoStopRecord(StopRecord)
	DoSetCurrentPreviewScene(SetCurrentPreviewScene)
	DoCreateScene(CreateScene)
	DoSetSceneName(SetSceneName)
	DoSetSceneSceneTransitionOverride(SetSceneSceneTransitionOverride)
	DoSetCurrentSceneTransition(SetCurrentSceneTransition)
	DoTriggerStudioModeTransition(TriggerStudioModeTransition)
	DoSetCurrentSceneCollection(SetCurrentSceneCollection)
	DoSetSceneItemBlendMode(SetSceneItemBlendMode)
	DoSetSceneItemTransform(SetSceneItemTransform)
	DoRemoveSceneItem(RemoveSceneItem)
	DoToggleVirtualCam(ToggleVirtualCam)
	DoStopVirtualCam(StopVirtualCam)
	DoStartVirtualCam(StartVirtualCam)
	DoToggleStream(ToggleStream)
	DoStartStream(StartStream)
	DoStopStream(StopStream)
	DoSendStreamCaption(SendStreamCaption)
	DoPauseRecord(PauseRecord)
	DoResumeRecord(ResumeRecord)
	DoToggleRecord(ToggleRecord)
	DoToggleRecordPause(ToggleRecordPause)

	/*OBS DO GET Manager Inits
	DoGetGroupList(GetGroupList)
	DoGetCurrentPreviewScene(GetCurrentPreviewScene)
	DoGetSceneSceneTransitionOverride(GetSceneSceneTransitionOverride)
	DoGetTransitionKindList(GetTransitionKindList)
	DoGetCurrentSceneTransition(GetCurrentSceneTransition)
	DoGetSceneCollectionList(GetSceneCollectionList)
	DoGetSceneItemBlendMode(GetSceneItemBlendMode)
	DoGetSceneItemEnabled(GetSceneItemEnabled)
	DoGetSceneItemTransform(GetSceneItemTransform)
	DoGetVirtualCamStatus(GetVirtualCamStatus)
	DoGetStreamStatus(GetStreamStatus)
	DoGetRecordStatus(GetRecordStatus)
	*/
}

type manager struct {
	logger *zap.Logger
	client *goobs.Client
}

func (m *manager) Close() error {
	_ = m.client.Disconnect()
	return nil
}

func (m *manager) DoSetScene(scene SetScene) {
	_, _ = m.client.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{SceneName: scene.SceneName})
}

func (m *manager) DoStartRecord(cmd StartRecord) {
	m.client.Record.StartRecord()
}

func (m *manager) DoStopRecord(cmd StopRecord) {
	m.client.Record.StopRecord()
}

//-------------------------------------------------------------

func (m *manager) DoSetCurrentPreviewScene(scene SetCurrentPreviewScene) {
	m.client.Scenes.SetCurrentPreviewScene(&scenes.SetCurrentPreviewSceneParams{SceneName: scene.PreviewSceneName})
}

func (m *manager) DoCreateScene(scene CreateScene) {
	m.client.Scenes.CreateScene(&scenes.CreateSceneParams{SceneName: scene.SceneName})
}

func (m *manager) DoSetSceneName(scene SetSceneName) {
	m.client.Scenes.SetSceneName(&scenes.SetSceneNameParams{SceneName: scene.SceneName})
}

func (m *manager) DoSetSceneSceneTransitionOverride(cmd SetSceneSceneTransitionOverride) {
	m.client.Scenes.SetSceneSceneTransitionOverride(&scenes.SetSceneSceneTransitionOverrideParams{
		SceneName:          cmd.SceneName,
		TransitionDuration: cmd.TransitionDuration,
		TransitionName:     cmd.TransitionName})
}

func (m *manager) DoSetCurrentSceneTransition(cmd SetCurrentSceneTransition) {
	m.client.Transitions.SetCurrentSceneTransition(&transitions.SetCurrentSceneTransitionParams{
		TransitionName: cmd.TransitionName})
}

func (m *manager) DoTriggerStudioModeTransition(cmd TriggerStudioModeTransition) {
	m.client.Transitions.TriggerStudioModeTransition(&transitions.TriggerStudioModeTransitionParams{})
}

func (m *manager) DoSetCurrentSceneCollection(scene SetCurrentSceneCollection) {
	m.client.Config.SetCurrentSceneCollection(&config.SetCurrentSceneCollectionParams{
		SceneCollectionName: scene.SceneCollectionName})
}

func (m *manager) DoSetSceneItemBlendMode(scene SetSceneItemBlendMode) {
	m.client.SceneItems.SetSceneItemBlendMode(&sceneitems.SetSceneItemBlendModeParams{
		SceneItemBlendMode: scene.SceneItemBlendMode,
		SceneItemId:        scene.SceneItemId,
		SceneName:          scene.SceneName})
}

func (m *manager) DoSetSceneItemTransform(sceneItem SetSceneItemTransform) {
	m.client.SceneItems.SetSceneItemTransform(&sceneitems.SetSceneItemTransformParams{
		SceneItemId: sceneItem.SceneItemId,
		SceneName:   sceneItem.SceneName,
		SceneItemTransform: &typedefs.SceneItemTransform{
			Alignment:       sceneItem.Alignment,
			BoundsAlignment: sceneItem.BoundsAlignment,
			BoundsHeight:    sceneItem.BoundsHeight,
			BoundsType:      sceneItem.BoundsType,
			BoundsWidth:     sceneItem.BoundsWidth,
			CropBottom:      sceneItem.CropBottom,
			CropLeft:        sceneItem.CropLeft,
			CropRight:       sceneItem.CropRight,
			CropTop:         sceneItem.CropTop,
			Height:          sceneItem.Height,
			PositionX:       sceneItem.PositionX,
			PositionY:       sceneItem.PositionY,
			Rotation:        sceneItem.Rotation,
			ScaleX:          sceneItem.ScaleX,
			ScaleY:          sceneItem.ScaleY,
			SourceHeight:    sceneItem.SourceHeight,
			SourceWidth:     sceneItem.SourceWidth,
			Width:           sceneItem.Width,
		},
	})
}

func (m *manager) DoRemoveSceneItem(sceneItem RemoveSceneItem) {
	m.client.SceneItems.RemoveSceneItem(&sceneitems.RemoveSceneItemParams{
		sceneItem.SceneItemId,
		sceneItem.SceneName,
	})
}

func (m *manager) DoToggleVirtualCam(cmd ToggleVirtualCam) {
	m.client.Outputs.ToggleVirtualCam()
}

func (m *manager) DoStopVirtualCam(cmd StopVirtualCam) {
	m.client.Outputs.StopVirtualCam()
}

func (m *manager) DoStartVirtualCam(cmd StartVirtualCam) {
	m.client.Outputs.StartVirtualCam()
}

func (m *manager) DoToggleStream(cmd ToggleStream) {
	m.client.Stream.ToggleStream()
}

func (m *manager) DoStartStream(cmd StartStream) {
	m.client.Stream.StartStream()
}

func (m *manager) DoStopStream(cmd StopStream) {
	m.client.Stream.StopStream()
}

func (m *manager) DoSendStreamCaption(cmd SendStreamCaption) {
	m.client.Stream.SendStreamCaption(&stream.SendStreamCaptionParams{CaptionText: cmd.StreamCaption})
}

func (m *manager) DoPauseRecord(cmd PauseRecord) {
	m.client.Record.PauseRecord()
}

func (m *manager) DoResumeRecord(cmd ResumeRecord) {
	m.client.Record.ResumeRecord()
}

func (m *manager) DoToggleRecord(cmd ToggleRecord) {
	m.client.Record.ToggleRecord()
}

func (m *manager) DoToggleRecordPause(cmd ToggleRecordPause) {
	m.client.Record.ToggleRecordPause()
}

/* OBS DO GET Methods
func (m *manager) DoGetSceneList(cmd GetSceneList) {
	m.client.Scenes.GetSceneList()
}

func (m *manager) DoGetGroupList(cmd GetGroupList) {
	m.client.Scenes.GetGroupList()
}

func (m *manager) DoGetCurrentPreviewScene(cmd GetCurrentPreviewScene) {
	m.client.Scenes.GetCurrentPreviewScene()
}
func (m *manager) DoGetSceneSceneTransitionOverride(cmd GetSceneSceneTransitionOverride) {
	//
}

func (m *manager) DoGetTransitionKindList(cmd GetTransitionKindList) {
	//
}

func (m *manager) DoGetCurrentSceneTransition(cmd GetCurrentSceneTransition) {
	//
}

func (m *manager) DoGetSceneCollectionList(cmd GetSceneCollectionList) {
	//
}

func (m *manager) DoGetSceneItemBlendMode(cmd GetSceneItemBlendMode) {
	//
}

func (m *manager) DoGetSceneItemEnabled(cmd GetSceneItemEnabled) {
	//
}

func (m *manager) DoGetSceneItemTransform(cmd GetSceneItemTransform) {
	//
}

func (m *manager) DoGetVirtualCamStatus(cmd GetVirtualCamStatus) {
	//
}

func (m *manager) DoGetStreamStatus(cmd GetStreamStatus) {
	//
}

func (m *manager) DoGetRecordStatus(cmd GetRecordStatus) {
	//
}
*/
//-------------------------------------------------------------

func NewManager(c ObsConf, logger *zap.Logger) (*manager, error) {
	client, err := goobs.New(c.HostPort, goobs.WithPassword(c.Password))
	return &manager{client: client, logger: logger.Named("OBSManager")}, err
}
