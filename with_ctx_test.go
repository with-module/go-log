package logger

import (
	"context"
	"testing"
)

type (
	RequestData struct {
		Flag      bool
		SecretKey string
	}
)

func (req RequestData) Customize() interface{} {
	ref := req
	if ref.SecretKey != "" {
		ref.SecretKey = "_hidden_"
	}

	return ref
}

func TestContextLogger(t *testing.T) {
	var getReqID GetRequestIDFromContextFunction = func(ctx context.Context) string {
		return "test_request_id"
	}

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

	ImportGetRequestIDFunction(getReqID)
	KInfo(context.Background(), "log context with level: %s", InfoLevel)
	KDError(context.Background(), RequestData{Flag: false, SecretKey: "very-secret"}, "log context with level: %s", ErrorLevel)
}
