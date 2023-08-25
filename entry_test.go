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

var defaultConfig = Config{
	Level:  "debug",
	Format: core.JsonFormat,
	Writer: WriterConfig{
		Output: []string{WriterConsole},
		Error:  []string{WriterConsoleError},
	},
}

func TestInitLogger(t *testing.T) {
	t.Run("fallback log level", func(t *testing.T) {
		err := LoadConfig(Config{
			Level:  "invalid-config",
			Format: core.JsonFormat,
			Writer: WriterConfig{
				Output: []string{WriterConsole},
				Error:  []string{WriterConsoleError},
			},
		})
		assert.NoError(t, err)
		assert.NotNil(t, inst)
		defer Flush()
		assert.Equal(t, DefaultLogLevel, inst.Level().String(), "unexpected log level")
		Info("print info log message")
	})

	t.Run("panic on invalid config", func(t *testing.T) {
		invalidConfig := Config{
			Level:  DefaultLogLevel,
			Format: core.PlainTextFormat,
			Writer: WriterConfig{
				Output: []string{"///invalid", "\\\\\\\\%invalid"},
				Error:  []string{"///invalid", "\\\\\\\\%invalid"},
			},
		}
		err := LoadConfig(invalidConfig)
		assert.ErrorContains(t, err, "failed to initiate default logger")
	})
}

func TestBasicLog(t *testing.T) {
	err := LoadConfig(defaultConfig)
	assert.NoError(t, err)
	assert.NotNil(t, inst)
	defer Flush()
	PrintLog("debug", "using global instance to print log with debug mode")
	Debug("print Debug log using default logger instance")
	Info("print Info log using default logger instance", param.Str("format", core.JsonFormat))
	Warn("print Warn log using default logger instance")
	testLog := Module("log-test")
	testLog.PrintLog("error", "this log will print at error level")
	testLog.Debugf("print debug message at %s", time.Now().Format(time.RFC3339))
}

func TestLogger_PrintLog(t *testing.T) {
	err := LoadConfig(defaultConfig)
	assert.NoError(t, err)
	assert.NotNil(t, inst)
	defer Flush()
	Error("print error message", param.Err(errors.New("invalid input error")))
}

func TestPanic(t *testing.T) {
	err := LoadConfig(defaultConfig)
	assert.NoError(t, err)
	assert.NotNil(t, inst)
	defer Flush()

	assert.Panics(t, func() {
		Panic("print this message and panic")
	})

	assert.Panics(t, func() {
		PrintLog("panic", "log message and panic")
	})
}

func TestFatal(t *testing.T) {
	err := LoadConfig(defaultConfig, zap.WithFatalHook(zapcore.WriteThenPanic))
	assert.NoError(t, err)
	assert.NotNil(t, inst)
	defer Flush()

	assert.Panics(t, func() {
		Fatal("print this message and panic")
	})

	assert.Panics(t, func() {
		PrintLog("fatal", "log message and panic")
	})
}
