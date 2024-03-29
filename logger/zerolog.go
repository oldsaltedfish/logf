package logger

import (
	"fmt"
	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	logger *zerolog.Logger
}

func NewZeroLogger(logger *zerolog.Logger) *ZeroLogger {
	return &ZeroLogger{logger: logger}
}

func (z ZeroLogger) Debug(i interface{}) {
	z.logger.Debug().Msg(fmt.Sprintf("%v", i))
}

func (z ZeroLogger) Info(i interface{}) {
	z.logger.Info().Msg(fmt.Sprintf("%v", i))
}

func (z ZeroLogger) Warn(i interface{}) {
	z.logger.Warn().Msg(fmt.Sprintf("%v", i))
}

func (z ZeroLogger) Error(i interface{}) {
	z.logger.Error().Msg(fmt.Sprintf("%v", i))
}

func (z ZeroLogger) Debugf(format string, a ...interface{}) {
	z.logger.Debug().Msg(fmt.Sprintf(format, a...))
}

func (z ZeroLogger) Infof(format string, a ...interface{}) {
	z.logger.Info().Msg(fmt.Sprintf(format, a...))
}

func (z ZeroLogger) Warnf(format string, a ...interface{}) {
	z.logger.Warn().Msg(fmt.Sprintf(format, a...))
}

func (z ZeroLogger) Errorf(format string, a ...interface{}) {
	z.logger.Error().Msg(fmt.Sprintf(format, a...))
}
