package core

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	*zap.SugaredLogger
}

func (c *Logger) Module(name string) *Logger {
	return &Logger{c.Named(name)}
}

func (c *Logger) PrintLog(mode string, msg string, args ...any) {
	level, err := zapcore.ParseLevel(mode)
	if err != nil {
		level = c.Level()
	}
	var fn func(msg string, args ...any)
	switch level {
	case zapcore.DebugLevel:
		fn = c.Debugw
	case zapcore.InfoLevel:
		fn = c.Infow
	case zapcore.ErrorLevel:
		fn = c.Errorw
	case zapcore.WarnLevel:
		fn = c.Warnw
	case zapcore.FatalLevel:
		fn = c.Fatalw
	default:
		fn = c.Panicw
	}
	fn(msg, args...)
}

func (c *Logger) Flush() error {
	err := c.Sync()
	if err == nil || errors.Is(err, os.ErrInvalid) {
		return nil
	}
	return err
}
