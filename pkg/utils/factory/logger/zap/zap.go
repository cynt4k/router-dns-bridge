package zap

import (
	"github.com/cynt4k/router-dns-bridge/cmd/config"
	loggerInterface "github.com/cynt4k/router-dns-bridge/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	rawZap *zap.Logger
}

func New(c *config.Config) loggerInterface.Logger {
	var level zap.AtomicLevel

	if c.DevMode {
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	} else {
		level = zap.NewAtomicLevel()
	}
	zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg := zap.Config{
		Level:       level,
		Development: c.DevMode,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	rawZap, _ := cfg.Build()
	return logger{
		rawZap: rawZap,
	}
}

func (z logger) New(name string) loggerInterface.Logger {
	newZap := z.rawZap.Named(name)
	return logger{
		rawZap: newZap,
	}
}
