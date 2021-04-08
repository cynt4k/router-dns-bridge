package logger

type mockLogger struct{}

func (m mockLogger) New(name string) Logger {
	return mockLogger{}
}

func (m mockLogger) Info(args ...interface{})                       {}
func (m mockLogger) Infof(template string, args ...interface{})     {}
func (m mockLogger) Infow(msg string, keysAndValues ...interface{}) {}

func (m mockLogger) Warn(args ...interface{})                       {}
func (m mockLogger) Warnf(template string, args ...interface{})     {}
func (m mockLogger) Warnw(msg string, keysAndValues ...interface{}) {}

func (m mockLogger) Debug(args ...interface{})                       {}
func (m mockLogger) Debugf(template string, args ...interface{})     {}
func (m mockLogger) Debugw(msg string, keysAndValues ...interface{}) {}

func (m mockLogger) Error(args ...interface{})                       {}
func (m mockLogger) Errorf(template string, args ...interface{})     {}
func (m mockLogger) Errorw(msg string, keysAndValues ...interface{}) {}

func (m mockLogger) Panic(args ...interface{})                       {}
func (m mockLogger) Panicf(template string, args ...interface{})     {}
func (m mockLogger) Panicw(msg string, keysAndValues ...interface{}) {}
