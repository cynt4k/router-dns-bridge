package logger

import "fmt"

var (
	logger Logger
)

type Logger interface {
	New(name string) Logger

	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})

	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
}

func SetLogger(logInstance Logger) {
	logger = logInstance
}

func Global() Logger {
	if logger == nil {
		fmt.Println("no global logger defined - no output will be displayed")
		logger = mockLogger{}
	}
	return logger
}
