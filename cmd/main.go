package main

import (
	"encoding/json"
	"flag"
	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	"git.miem.hse.ru/hubman/hubman-lib/executor"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"obs-man/internal"
	"os"
	"time"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "config-path", "conf.yaml", "specify config path")
	flag.Parse()

	c := obsman.Conf{}
	f, err := os.Open(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	obsManager, err := obsman.NewManager(c.ObsConf, zap.Must(zap.NewProduction()))

	if err != nil {
		log.Fatal(err)
	}

	app := hubman.NewAgentApp(
		core.AgentConfiguration{
			System: &core.SystemConfig{
				Server: &core.InterfaceConfig{
					IP:   net.IPv4(127, 0, 0, 1),
					Port: c.Agent.ExecutorPort,
				},
				RedisUrl: c.Agent.RedisUrl,
			},
			User: &c.ObsConf,
			ParseUserConfig: func(jsonBytes []byte) (core.Configuration, error) {
				newConf := obsman.ObsConf{}
				err := json.Unmarshal(jsonBytes, &newConf)
				return &newConf, err
			},
		},
		hubman.WithExecutor(
			hubman.WithCommand(obsman.SetScene{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.SetScene{}
				parser(&cmd)
				obsManager.DoSetScene(cmd)
			}),
			hubman.WithCommand(obsman.StartRecord{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.StartRecord{}
				parser(&cmd)
				obsManager.DoStartRecord(cmd)
			}),
			hubman.WithCommand(obsman.StopRecord{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.StopRecord{}
				parser(&cmd)
				obsManager.DoStopRecord(cmd)
			}),
			hubman.WithCommand(obsman.SetCurrentPreviewScene{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.SetCurrentPreviewScene{}
				parser(&cmd)
				obsManager.DoSetCurrentPreviewScene(cmd)
			}),
			hubman.WithCommand(obsman.CreateScene{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.CreateScene{}
				parser(&cmd)
				obsManager.DoCreateScene(cmd)
			}),
			hubman.WithCommand(obsman.SetSceneName{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.SetSceneName{}
				parser(&cmd)
				obsManager.DoSetSceneName(cmd)
			}),
			hubman.WithCommand(obsman.SetSceneSceneTransitionOverride{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.SetSceneSceneTransitionOverride{}
				parser(&cmd)
				obsManager.DoSetSceneSceneTransitionOverride(cmd)
			}),
			hubman.WithCommand(obsman.SetCurrentSceneTransition{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.SetCurrentSceneTransition{}
				parser(&cmd)
				obsManager.DoSetCurrentSceneTransition(cmd)
			}),
			hubman.WithCommand(obsman.TriggerStudioModeTransition{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.TriggerStudioModeTransition{}
				parser(&cmd)
				obsManager.DoTriggerStudioModeTransition(cmd)
			}),
			hubman.WithCommand(obsman.SetCurrentSceneCollection{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.SetCurrentSceneCollection{}
				parser(&cmd)
				obsManager.DoSetCurrentSceneCollection(cmd)
			}),
			hubman.WithCommand(obsman.SetSceneItemBlendMode{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.SetSceneItemBlendMode{}
				parser(&cmd)
				obsManager.DoSetSceneItemBlendMode(cmd)
			}),
			hubman.WithCommand(obsman.SetSceneItemTransform{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.SetSceneItemTransform{}
				parser(&cmd)
				obsManager.DoSetSceneItemTransform(cmd)
			}),
			hubman.WithCommand(obsman.RemoveSceneItem{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.RemoveSceneItem{}
				parser(&cmd)
				obsManager.DoRemoveSceneItem(cmd)
			}),
			hubman.WithCommand(obsman.ToggleVirtualCam{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.ToggleVirtualCam{}
				parser(&cmd)
				obsManager.DoToggleVirtualCam(cmd)
			}),
			hubman.WithCommand(obsman.StopVirtualCam{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.StopVirtualCam{}
				parser(&cmd)
				obsManager.DoStopVirtualCam(cmd)
			}),
			hubman.WithCommand(obsman.StartVirtualCam{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.StartVirtualCam{}
				parser(&cmd)
				obsManager.DoStartVirtualCam(cmd)
			}),
			hubman.WithCommand(obsman.ToggleStream{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.ToggleStream{}
				parser(&cmd)
				obsManager.DoToggleStream(cmd)
			}),
			hubman.WithCommand(obsman.StartStream{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.StartStream{}
				parser(&cmd)
				obsManager.DoStartStream(cmd)
			}),
			hubman.WithCommand(obsman.StopStream{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.StopStream{}
				parser(&cmd)
				obsManager.DoStopStream(cmd)
			}),
			hubman.WithCommand(obsman.SendStreamCaption{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.SendStreamCaption{}
				parser(&cmd)
				obsManager.DoSendStreamCaption(cmd)
			}),
			hubman.WithCommand(obsman.PauseRecord{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.PauseRecord{}
				parser(&cmd)
				obsManager.DoPauseRecord(cmd)
			}),
			hubman.WithCommand(obsman.ResumeRecord{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.ResumeRecord{}
				parser(&cmd)
				obsManager.DoResumeRecord(cmd)
			}),
			hubman.WithCommand(obsman.StartRecord{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.StartRecord{}
				parser(&cmd)
				obsManager.DoStartRecord(cmd)
			}),
			hubman.WithCommand(obsman.StopRecord{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.StopRecord{}
				parser(&cmd)
				obsManager.DoStopRecord(cmd)
			}),
			hubman.WithCommand(obsman.ToggleRecord{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.ToggleRecord{}
				parser(&cmd)
				obsManager.DoToggleRecord(cmd)
			}),
			hubman.WithCommand(obsman.ToggleRecordPause{}, func(command core.SerializedCommand, parser executor.CommandParser) {
				cmd := obsman.ToggleRecordPause{}
				parser(&cmd)
				obsManager.DoToggleRecordPause(cmd)
			}),
		),
		hubman.WithOnConfigRefresh(func(c core.AgentConfiguration) {
			<-time.After(time.Second * 1)
		}),
	)

	<-app.WaitShutdown()
	obsManager.Close()
}
