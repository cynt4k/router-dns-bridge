package zap

func (z logger) Info(args ...interface{}) {
	z.rawZap.Sugar().Info(args...)
}

func (z logger) Infof(template string, args ...interface{}) {
	z.rawZap.Sugar().Infof(template, args...)
}

func (z logger) Infow(msg string, args ...interface{}) {
	z.rawZap.Sugar().Infow(msg, args...)
}

func (z logger) Warn(args ...interface{}) {
	z.rawZap.Sugar().Warn(args...)
}

func (z logger) Warnf(template string, args ...interface{}) {
	z.rawZap.Sugar().Warnf(template, args...)
}

func (z logger) Warnw(msg string, args ...interface{}) {
	z.rawZap.Sugar().Warnw(msg, args...)
}

func (z logger) Error(args ...interface{}) {
	z.rawZap.Sugar().Error(args...)
}

func (z logger) Errorf(template string, args ...interface{}) {
	z.rawZap.Sugar().Errorf(template, args...)
}

func (z logger) Errorw(msg string, args ...interface{}) {
	z.rawZap.Sugar().Errorw(msg, args...)
}

func (z logger) Debug(args ...interface{}) {
	z.rawZap.Sugar().Debug(args...)
}

func (z logger) Debugf(template string, args ...interface{}) {
	z.rawZap.Sugar().Debugf(template, args...)
}

func (z logger) Debugw(msg string, args ...interface{}) {
	z.rawZap.Sugar().Debugw(msg, args...)
}

func (z logger) Panic(args ...interface{}) {
	z.rawZap.Sugar().Panic(args...)
}

func (z logger) Panicf(template string, args ...interface{}) {
	z.rawZap.Sugar().Panicf(template, args...)
}

func (z logger) Panicw(msg string, args ...interface{}) {
	z.rawZap.Sugar().Panicw(msg, args...)
}
