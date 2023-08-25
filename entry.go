package log

import (
	"fmt"
	"gitlab.com/with-junbach/go-modules/log/core"
	"gitlab.com/with-junbach/go-modules/log/param"
	"go.uber.org/zap"
	"log"
)

const (
	WriterConsole      = "stdout"
	WriterConsoleError = "stderr"

	DefaultLogLevel = "info"
)

var inst *core.BaseLogger

type (
	Config       = core.Config
	WriterConfig = core.WriterConfig
)

func init() {
	if err := LoadConfig(Config{
		Level:  DefaultLogLevel,
		Format: core.PlainTextFormat,
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	}); err != nil {
		log.Panicf("error: %s", err)
	}
}

func LoadConfig(config Config, opts ...zap.Option) error {
	initLogger, err := core.InitLogger(config, opts...)
	if err != nil {
		return fmt.Errorf("failed to initiate default logger: %w", err)
	}
	inst = initLogger
	inst.Debugw("default logger instance has been initialized successfully", param.Obj("config", config))
	return nil
}

func Flush() {
	if err := inst.Close(); err != nil {
		inst.Errorw("failed to flush logger buffer", param.Err(err))
	}
}

func Module(name string) *core.BaseLogger {
	return inst.Module(name)
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
