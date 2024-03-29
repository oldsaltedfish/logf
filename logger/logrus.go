package logger

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger(l *logrus.Logger) *LogrusLogger {
	return &LogrusLogger{logger: l}
}

func (l *LogrusLogger) Debug(i interface{}) {
	l.logger.Debug(i)
}

func (l *LogrusLogger) Info(i interface{}) {
	l.logger.Debug(i)
}

func (l *LogrusLogger) Warn(i interface{}) {
	l.logger.Warn(i)
}

func (l *LogrusLogger) Error(i interface{}) {
	l.logger.Error(i)
}

func (l *LogrusLogger) Debugf(format string, a ...interface{}) {
	l.logger.Debugf(format, a...)
}

func (l *LogrusLogger) Infof(format string, a ...interface{}) {
	l.logger.Infof(format, a...)
}

func (l *LogrusLogger) Warnf(format string, a ...interface{}) {
	l.logger.Warnf(format, a...)
}

func (l *LogrusLogger) Errorf(format string, a ...interface{}) {
	l.logger.Errorf(format, a...)
}
