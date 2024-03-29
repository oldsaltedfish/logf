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

func NewLoggerWithZap() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger.NewZapLogger(l.Sugar())
}

func NewLoggerWithLogrus() {
	logger.NewLogrusLogger(logrus.New())
}

func NewLoggerWithZeroLog() {
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.NewZeroLogger(&l)
}
