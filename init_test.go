package log

import (
	"os"
	"testing"
)

func TestLoadConfig(withTest *testing.T) {

	withTest.Run("LoadValidConfig", func(t *testing.T) {
		cfg := Config{
			Module: "log-test-svc",
			Level:  "error",
			Output: "stderr",
		}
		LoadConfig(cfg, WithOutput(os.Stdout))
		if std.GetLevel() != ErrorLevel {
			t.Errorf("log level must be [error], got: %s", std.GetLevel())
		}
	})

	withTest.Run("LoadInvalidConfig", func(t *testing.T) {
		cfg := Config{
			Module: "not-a-valid-svc",
			Level:  "invalid-level",
			Output: "stderr",
		}
		LoadConfig(cfg)
		if std.GetLevel() != DebugLevel {
			t.Errorf("log level must fallback to [debug], got: %s", std.GetLevel())
		}
	})
}
