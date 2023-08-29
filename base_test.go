package log

import (
	"context"
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
	Format: core.FormatJson,
}

func TestInitLogger(t *testing.T) {
	t.Run("fallback log level", func(t *testing.T) {
		err := LoadConfig(Config{
			Level:  "invalid-config",
			Format: core.FormatJson,
		})
		assert.NoError(t, err)
		defer Flush()
		assert.Equal(t, zapcore.InfoLevel, inst.Level(), "unexpected log level")
		Info("print info log message")
	})

	t.Run("init logger failed due to invalid writers", func(t *testing.T) {
		invalidConfig := Config{
			Level:  "debug",
			Format: core.FormatJson,
			Writer: core.WriterConfig{
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
	defer Flush()
	PrintLog("debug", "using global instance to print log with debug mode")
	Debug("print Debug log using default logger instance")
	Info("print Info log using default logger instance", param.Str("format", core.FormatJson))
	Warn("print Warn log using default logger instance", param.Int("epoch", time.Now().Unix()))
	Error("print error message", param.Err(errors.New("invalid input error")))
	log := Module("log-test")
	assert.NotNil(t, log)
	assert.IsType(t, new(Logger), log)
	log.Errorln("this log will print at error level")
}

func TestPanic(t *testing.T) {
	err := LoadConfig(defaultConfig)
	assert.NoError(t, err)
	defer Flush()

	assert.Panics(t, func() {
		Panic("print this message and panic")
	})
}

func TestFatal(t *testing.T) {
	err := LoadConfig(defaultConfig, zap.WithFatalHook(zapcore.WriteThenPanic))
	assert.NoError(t, err)
	defer Flush()

	assert.Panics(t, func() {
		Fatal("print this message and panic")
	})
}

func TestCtxLog(t *testing.T) {
	err := LoadConfig(Config{
		Level:  "debug",
		Format: "json",
	})
	assert.NoError(t, err)
	defer Flush()

	ctx := BindContext(context.Background(), "request_id", "some-id")
	assert.Implements(t, (*context.Context)(nil), ctx)

	log := Ctx(ctx)
	assert.IsType(t, new(Logger), log)

	CtxLog(ctx, "invalid-mode", "this log will print at debug level and use global inst due to empty ctx")
	log.CtxLog(ctx, "info", "print info message with request-id and latency duration", param.Dur("responseTime", time.Millisecond*200))

}
