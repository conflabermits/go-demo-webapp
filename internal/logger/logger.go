package logger

import (
	"go.uber.org/zap"
)

func New(logLevel string) (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel) // Adjust log level as needed

	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return l, nil
}
