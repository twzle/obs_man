package command

import (
	"github.com/andreykaipov/goobs"
	"go.uber.org/zap"
)

// Representation of runnable command entity
type RunnableCommand interface {
	Run(ObsProvider, *zap.Logger) error
}

// Representation of handler provider entity
type ObsProvider interface {
	Provide() (*goobs.Client, error)
}

// Function logging error messages
func logErr(log *zap.Logger, msg string, err error) error {
	if err != nil {
		log.Error(msg, zap.Error(err))
	}
	return err
}
