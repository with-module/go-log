package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitLogger(t *testing.T) {
	t.Run("fallback log level", func(t *testing.T) {
		LoadConfig(Config{
			Level:  "invalid-config",
			Format: JsonFormat,
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
				Format: PlainTextFormat,
				Writer: WriterConfig{
					Output: []string{"///invalid"},
					Error:  []string{"///invalid"},
				},
			}
			LoadConfig(invalidConfig)
		}
		assert.Panics(t, fn, "must panic")
	})
}
