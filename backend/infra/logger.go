package infra

import "go.uber.org/zap"

type Logger struct {
	*zap.Logger
}

func MustNewLogger(cfg *Logging) *Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return &Logger{logger}
}
