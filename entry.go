package log

import (
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
	LoadConfig(Config{
		Level:  DefaultLogLevel,
		Format: core.PlainTextFormat,
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	})
}

func LoadConfig(config Config, opts ...zap.Option) {
	coreLogger, err := core.InitLogger(config, opts...)
	if err != nil {
		log.Panicf("failed to initiate default logger instance: %s", err)
	}

	inst = coreLogger
	inst.Debugw("default logger instance has been initialized successfully", param.Obj("config", config))
}

func Flush() {
	err := inst.Close()
	if err != nil {
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
