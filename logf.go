package logf

import (
	"github.com/oldsaltedfish/logf/logger"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"os"
)

type Logger interface {
	Debug(interface{})
	Info(interface{})
	Warn(interface{})
	Error(interface{})

	Debugf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
}

func NewLoggerWithZap() Logger {
	l, err := zap.NewProductionConfig().Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l.Sugar())
}

func NewLoggerWithLogrus() Logger {
	return logger.NewLogrusLogger(logrus.New())
}

func NewLoggerWithZeroLog() Logger {
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return logger.NewZeroLogger(&l)
}
