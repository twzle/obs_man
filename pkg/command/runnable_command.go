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

func logErr(log *zap.Logger, msg string, err error) error {
	if err != nil {
		log.Error(msg, zap.Error(err))
	}
	return err
}
