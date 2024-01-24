package main

import (
	"encoding/json"
	"log"
	obsman "obs-man/pkg"
	ocmd "obs-man/pkg/command"
	osig "obs-man/pkg/signal"

	"git.miem.hse.ru/hubman/hubman-lib"
	"git.miem.hse.ru/hubman/hubman-lib/core"
	hcore "git.miem.hse.ru/hubman/hubman-lib/core"
	"git.miem.hse.ru/hubman/hubman-lib/executor"
	"git.miem.hse.ru/hubman/hubman-lib/manipulator"
	"go.uber.org/zap"
)

func main() {
	systemConfig := &hcore.SystemConfig{}
	userConfig := &obsman.Conf{}

	err := hcore.ReadConfig(systemConfig, userConfig)
	if err != nil {
		log.Fatalf("error while init config: %v", err)
	}

	conf := hcore.AgentConfiguration{
		System: systemConfig,
		User:   userConfig,
		ParseUserConfig: func(jsonBuf []byte) (hcore.Configuration, error) {
			var conf *obsman.Conf
			err := json.Unmarshal(jsonBuf, &conf)
			return conf, err
		},
	}
	signalsCh := make(chan hcore.Signal)

	app := hcore.NewContainer(conf.System.Logging)

	obsManager, err := obsman.NewManager(userConfig.ObsConf, app.Logger(), signalsCh)
	if err != nil {
		app.Logger().Warn("Unable to create obs manager", zap.Error(err))
	}
	go obsManager.HealthCheck(userConfig.ObsConf, app.WaitShutdown())
	defer obsManager.Close()

	executorCommands := append(make([]func(executor.Executor), 0), ocmd.ProvideRecordCommands(obsManager, app.Logger())...)
	executorCommands = append(executorCommands, ocmd.ProvideSceneCommands(obsManager, app.Logger())...)
	executorCommands = append(executorCommands, ocmd.ProvideSourceCommands(obsManager, app.Logger())...)
	executorCommands = append(executorCommands, ocmd.ProvideStreamCommands(obsManager, app.Logger())...)
	executorCommands = append(executorCommands, ocmd.ProvideUiCommands(obsManager, app.Logger())...)
	executorCommands = append(executorCommands, ocmd.ProvideTransitionCommands(obsManager, app.Logger())...)

	manipulatorSignals := append(make([]func(manipulator.Manipulator), 0), hubman.WithChannel(signalsCh))
	manipulatorSignals = append(manipulatorSignals, osig.CurrentSceneSignals...)
	manipulatorSignals = append(manipulatorSignals, osig.InputSignals...)
	manipulatorSignals = append(manipulatorSignals, osig.RecordSignals...)
	manipulatorSignals = append(manipulatorSignals, osig.ReplayBufferSignals...)
	manipulatorSignals = append(manipulatorSignals, osig.SceneItemSignals...)
	manipulatorSignals = append(manipulatorSignals, osig.ScreenshotSignals...)
	manipulatorSignals = append(manipulatorSignals, osig.StreamSignals...)
	manipulatorSignals = append(manipulatorSignals, osig.TransitionSignals...)
	manipulatorSignals = append(manipulatorSignals, osig.UISignals...)
	manipulatorSignals = append(manipulatorSignals, osig.VirtualCamSignals...)

	app.RegisterPlugin(hubman.NewAgentPlugin(app.Logger(),
		conf,
		hubman.WithExecutor(
			executorCommands...,
		),
		hubman.WithManipulator(
			manipulatorSignals...,
		),
		hubman.WithOnConfigRefresh(func(c core.AgentConfiguration) {
			newConf, ok := c.User.(*obsman.Conf)
			if !ok {
				log.Fatal("Invalid new config provided", zap.Any("newConf", c.User))
			}
			obsManager.UpdateConn(newConf.ObsConf)
		}),
	))

	<-app.WaitShutdown()
}
