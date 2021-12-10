package log

import (
	"fmt"
	"net/http"
	"testing"
)

func TestEntryLogger(t *testing.T) {
	testConfig := Config{
		Level:     "debug",
		Output:    []string{"stdout"},
		ErrOutput: []string{"stdout", "stderr"},
		Format:    "console",
		Module:    "test-module",
	}

	err := InitLoggerInst(testConfig)
	if err != nil {
		t.Fatalf("should not have this error: %v", err)
	}
	defer Close()

	Info("Print sample info test")
	Print(ErrorLevel, "sample error message with custom fields", "field_name", "test")
	fields := make(map[string]interface{})
	fields["request_id"] = "1234-5678-abcd"
	fields["protocol"] = "http"
	fields["user_agent"] = "postman"
	fields["remote_ip"] = "::1"
	fields["status"] = http.StatusOK
	PrintMap(InfoLevel, fmt.Sprintf("request returned with status: %d", fields["status"]), fields)
}
