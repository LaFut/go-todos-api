package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	logrus.Logger
}

func NewLogger() *Logger {
	logger := &Logger{
		logrus.Logger{
			Out:       os.Stdout,
			Formatter: &logrus.TextFormatter{ForceColors: true},
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.InfoLevel,
		},
	}

	return logger
}
