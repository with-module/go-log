package log

import (
	"os"
	"testing"
)

func TestLoadConfig(withTest *testing.T) {

	withTest.Run("LoadValidConfig", func(t *testing.T) {
		cfg := Config{
			Module: "log-test-svc",
			Level:  "debug",
			Output: "stderr",
		}
		err := LoadConfig(cfg, WithOutput(os.Stdout))
		if err != nil {
			t.Errorf("load log config must not fail: %v", err)
		}
	})

	withTest.Run("LoadInvalidConfig", func(t *testing.T) {
		cfg := Config{
			Module: "not-a-valid-svc",
			Level:  "invalid-level",
			Output: "stderr",
		}
		err := LoadConfig(cfg)
		Error(err, "got error when load log config")
		if err == nil {
			t.Errorf("there must be error due to invalid config")
		}
	})
}
