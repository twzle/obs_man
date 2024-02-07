package command

import (
	"git.miem.hse.ru/hubman/hubman-lib/core"
	"github.com/andreykaipov/goobs"
	"go.uber.org/zap"
)

type RunnableCommand interface {
	Run(ObsProvider, *zap.Logger) error
}

type ObsProvider interface {
	Provide() (*goobs.Client, chan<- core.Signal, error)
}
