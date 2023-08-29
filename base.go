package log

import (
	"context"
	"fmt"
	"gitlab.com/with-junbach/go-modules/log/core"
	"gitlab.com/with-junbach/go-modules/log/param"
	"log"
)

type (
	Config = core.Config
	Logger = core.Logger
)

var inst *Logger

func init() {
	if err := LoadConfig(Config{
		Level:  "debug",
		Format: core.FormatJson,
	}); err != nil {
		log.Panicf("error: %s", err)
	}
}

func LoadConfig(config Config, opts ...core.Option) error {
	initLogger, err := core.NewLogger(config, opts...)
	if err != nil {
		return fmt.Errorf("failed to initiate default logger: %w", err)
	}
	inst = initLogger
	inst.Debugw("default logger instance has been initialized successfully", param.Any("config", config))
	return nil
}

func Flush() {
	inst.Debug("flush logger buffer")
	if err := inst.Flush(); err != nil {
		inst.Errorw("failed to flush logger buffer", param.Err(err))
	}
}

func Module(name string, opts ...core.Option) *Logger {
	return inst.Module(name, opts...)
}

func PrintLog(mode string, msg string, args ...any) {
	inst.PrintLog(mode, msg, args...)
}

func Debug(msg string, args ...any) {
	inst.Debugw(msg, args...)
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

func BindContext(ctx context.Context, fields ...any) context.Context {
	return inst.BindContext(ctx, fields...)
}

func Ctx(ctx context.Context) *Logger {
	return inst.Ctx(ctx)
}

func CtxLog(ctx context.Context, mode string, msg string, args ...any) {
	inst.CtxLog(ctx, mode, msg, args...)
}
