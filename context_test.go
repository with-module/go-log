package log

import (
	"context"
	"testing"
	"time"
)

func TestCtxLog(t *testing.T) {
	LoadConfig(Config{
		Level:  "debug",
		Format: "json",
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	})
	defer Flush()
	CtxLog(context.Background(), "invalid-mode", "this log will print at debug level and use global inst due to empty ctx")
	log := Module("log-context")
	ctx := BindContext(context.Background(), "request_id", "some-id")
	log.CtxLog(ctx, "info", "print info message with request-id and latency duration", "latency", time.Millisecond*200)
}
