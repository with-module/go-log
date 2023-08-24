package log

import (
	"context"
	"testing"
	"time"
)

func TestBasicLog(t *testing.T) {
	LoadConfig(Config{
		Level:       "debug",
		Format:      "json",
		Output:      []string{ConsoleOutput},
		ErrorOutput: []string{ConsoleOutputError},
	})
	PrintLog("debug", "using global instance to print log with debug mode")
	CtxLog(context.Background(), "invalid-mode", "this log will print at debug level and use global inst due to empty ctx")
	log := WithModule("log-test")
	log.PrintLog("error", "this log will print at error level")
	log.Debugf("print debug message at %s", time.Now().Format(time.RFC3339))
	ctx := BindContext(context.Background(), "request_id", "some-id")
	log.Ctx(ctx).PrintLog("info", "print info message with request-id and latency duration", "latency", time.Millisecond*200)
}
