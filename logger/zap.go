package logger

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger(sugar *zap.SugaredLogger) *ZapLogger {
	return &ZapLogger{logger: sugar}
}

func (l *ZapLogger) Debugf(format string, a ...interface{}) {
	l.logger.Debugf(format, a...)
}

func (l *ZapLogger) Infof(format string, a ...interface{}) {
	l.logger.Infof(format, a...)
}

func (l *ZapLogger) Warnf(format string, a ...interface{}) {
	l.logger.Warnf(format, a...)
}

func (l *ZapLogger) Errorf(format string, a ...interface{}) {
	l.logger.Errorf(format, a...)
}

func (l *ZapLogger) Debug(i interface{}) {
	l.logger.Debug(i)
}

func (l *ZapLogger) Info(i interface{}) {
	l.logger.Info(i)
}

func (l *ZapLogger) Warn(i interface{}) {
	l.logger.Warn(i)
}

func (l *ZapLogger) Error(i interface{}) {
	l.logger.Error(i)
}
