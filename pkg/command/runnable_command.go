package command

import (
	"github.com/andreykaipov/goobs"
	"go.uber.org/zap"
	"git.miem.hse.ru/hubman/hubman-lib/core"
)

type RunnableCommand interface {
	Run(ObsProvider, *zap.Logger) error
}

type ObsProvider interface {
	Provide() (*goobs.Client, chan<- core.Signal, error)
}
