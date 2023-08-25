package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	BaseLogger struct {
		*zap.SugaredLogger
	}

	Printer func(msg string, args ...any)
)

func (b *BaseLogger) Module(name string) *BaseLogger {
	return &BaseLogger{b.Named(name)}
}

func (b *BaseLogger) PrintLog(mode string, msg string, args ...any) {
	level, err := zapcore.ParseLevel(mode)
	if err != nil {
		level = b.Level()
	}

	var fn Printer
	switch level {
	case zapcore.DebugLevel:
		fn = b.Debugw
	case zapcore.InfoLevel:
		fn = b.Infow
	case zapcore.ErrorLevel:
		fn = b.Errorw
	case zapcore.FatalLevel:
		fn = b.Fatalw
	default:
		fn = b.Panicw
	}

	fn(msg, args...)
}

func (b *BaseLogger) Close() error {
	return b.Sync()
}
