package log

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func TestBasicLog(t *testing.T) {
	config := Config{
		Level:  "debug",
		Format: JsonFormat,
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	}
	LoadConfig(config)
	defer Flush()
	PrintLog("debug", "using global instance to print log with debug mode")
	Debug("print Debug log using default logger instance")
	Info("print Info log using default logger instance", AddString("format", config.Format))
	Warn("print Warn log using default logger instance")
	log := WithModule("log-test")
	log.PrintLog("error", "this log will print at error level")
	log.Debugf("print debug message at %s", time.Now().Format(time.RFC3339))
}

func TestLogger_PrintLog(t *testing.T) {
	LoadConfig(Config{
		Level:  "debug",
		Format: JsonFormat,
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	})
	defer Flush()
	Error("print error message", AddError(errors.New("invalid input error")))
}

func TestPanic(t *testing.T) {
	LoadConfig(Config{
		Level:  "debug",
		Format: JsonFormat,
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	})
	defer Flush()

	assert.Panics(t, func() {
		Panic("print this message and panic")
	})

	assert.Panics(t, func() {
		PrintLog("panic", "log message and panic")
	})
}

func TestFatal(t *testing.T) {
	LoadConfig(Config{
		Level:  "debug",
		Format: JsonFormat,
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	}, zap.WithFatalHook(zapcore.WriteThenPanic))
	defer Flush()

	assert.Panics(t, func() {
		Fatal("print this message and panic")
	})

	assert.Panics(t, func() {
		PrintLog("fatal", "log message and panic")
	})
}
