package log

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gitlab.com/with-junbach/go-modules/log/core"
	"gitlab.com/with-junbach/go-modules/log/param"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func TestInitLogger(t *testing.T) {
	t.Run("fallback log level", func(t *testing.T) {
		LoadConfig(Config{
			Level:  "invalid-config",
			Format: core.JsonFormat,
			Writer: WriterConfig{
				Output: []string{WriterConsole},
				Error:  []string{WriterConsoleError},
			},
		})
		defer Flush()
		assert.Equal(t, DefaultLogLevel, inst.Level().String(), "unexpected log level")
		Info("print info log message")
	})

	t.Run("panic on invalid config", func(t *testing.T) {
		fn := func() {
			invalidConfig := Config{
				Level:  DefaultLogLevel,
				Format: core.PlainTextFormat,
				Writer: WriterConfig{
					Output: []string{"///invalid", "\\\\\\\\%invalid"},
					Error:  []string{"///invalid", "\\\\\\\\%invalid"},
				},
			}
			LoadConfig(invalidConfig)
		}
		assert.Panics(t, fn, "must panic")
	})
}

func TestBasicLog(t *testing.T) {
	config := Config{
		Level:  "debug",
		Format: core.JsonFormat,
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	}
	LoadConfig(config)
	defer Flush()
	PrintLog("debug", "using global instance to print log with debug mode")
	Debug("print Debug log using default logger instance")
	Info("print Info log using default logger instance", param.Str("format", config.Format))
	Warn("print Warn log using default logger instance")
	log := Module("log-test")
	log.PrintLog("error", "this log will print at error level")
	log.Debugf("print debug message at %s", time.Now().Format(time.RFC3339))
}

func TestLogger_PrintLog(t *testing.T) {
	LoadConfig(Config{
		Level:  "debug",
		Format: core.JsonFormat,
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	})
	defer Flush()
	Error("print error message", param.Err(errors.New("invalid input error")))
}

func TestPanic(t *testing.T) {
	LoadConfig(Config{
		Level:  "debug",
		Format: core.JsonFormat,
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
		Format: core.JsonFormat,
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
