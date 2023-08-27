package core

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"slices"
	"syscall"
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

func isAcceptedFlushErr(inputErr error) bool {
	return slices.ContainsFunc([]error{os.ErrInvalid, syscall.EBADF, syscall.ENOTTY}, func(err error) bool {
		return errors.Is(inputErr, err)
	})
}

func (c *Logger) Flush() error {
	err := c.Sync()
	if err == nil || isAcceptedFlushErr(err) {
		return nil
	}
	return err
}

func (c *Logger) Debug(msg string, args ...any) {
	c.Debugw(msg, args...)
}

func (c *Logger) Info(msg string, args ...any) {
	c.Infow(msg, args...)
}

func (c *Logger) Warn(msg string, args ...any) {
	c.Warnw(msg, args...)
}

func (c *Logger) Error(msg string, args ...any) {
	c.Errorw(msg, args...)
}

func (c *Logger) Fatal(msg string, args ...any) {
	c.Fatalw(msg, args...)
}

func (c *Logger) Panic(msg string, args ...any) {
	c.Panicw(msg, args...)
}
