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
		),
		hubman.WithOnConfigRefresh(func(c core.AgentConfiguration) {
			<-time.After(time.Second * 1)
		}),
	)

	<-app.WaitShutdown()
	obsManager.Close()
}
