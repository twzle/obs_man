package command

import (
	"github.com/andreykaipov/goobs"
	"go.uber.org/zap"
)

type RunnableCommand interface {
	Run(ObsProvider, *zap.Logger) error
}

type ObsProvider interface {
	Provide() (*goobs.Client, error)
}
