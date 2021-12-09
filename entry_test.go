package logger

import "testing"

func TestEntryLogger(t *testing.T) {
	testConfig := Config{
		Level:     "debug",
		Output:    []string{"stdout"},
		ErrOutput: []string{"stdout", "stderr"},
		Format:    "json",
		Module:    "test-module",
	}

	err := InitLoggerInst(testConfig)
	if err != nil {
		t.Fatalf("should not have this error: %v", err)
	}

	Info("Print sample info test")
	Print(ErrorLevel, "sample error message with custom fields", "field_name", "test")
}
