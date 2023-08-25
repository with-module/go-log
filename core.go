package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Logger struct {
		*zap.SugaredLogger
	}

	Printer func(msg string, args ...any)
)

var inst *Logger

func WithModule(name string) *Logger {
	return inst.WithModule(name)
}

func PrintLog(mode string, msg string, args ...any) {
	inst.PrintLog(mode, msg, args...)
}

func (c *Logger) WithModule(name string) *Logger {
	return &Logger{c.Named(name)}
}

func (c *Logger) PrintLog(mode string, msg string, args ...any) {
	level, err := zapcore.ParseLevel(mode)
	if err != nil {
		level = c.Level()
	}

	var fn Printer
	switch level {
	case zapcore.DebugLevel:
		fn = c.Debugw
	case zapcore.InfoLevel:
		fn = c.Infow
	case zapcore.ErrorLevel:
		fn = c.Errorw
	case zapcore.FatalLevel:
		fn = c.Fatalw
	default:
		fn = c.Panicw
	}

	fn(msg, args...)
}

func Debug(msg string, args ...any) {
	inst.Infow(msg, args...)
}

func Info(msg string, args ...any) {
	inst.Infow(msg, args...)
}

func Warn(msg string, args ...any) {
	inst.Warnw(msg, args...)
}

func Error(msg string, args ...any) {
	inst.Errorw(msg, args...)
}

func Fatal(msg string, args ...any) {
	inst.Fatalw(msg, args...)
}

func Panic(msg string, args ...any) {
	inst.Panicw(msg, args...)
}
