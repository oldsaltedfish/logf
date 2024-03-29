package logf

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
