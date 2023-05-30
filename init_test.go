package log

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	cfg := Config{
		Module: "not-a-valid-svc",
		Level:  "invalid-level",
		Output: "stderr",
		Caller: struct {
			Enabled   bool `config:"Enabled"`
			SkipFrame int  `config:"SkipFrame"`
		}{Enabled: true, SkipFrame: 2},
	}
	LoadConfig(cfg, WithOutput(os.Stdout))
	if std.GetLevel() != DebugLevel {
		t.Errorf("log level must fallback to [debug], got: %s", std.GetLevel())
	}
}
