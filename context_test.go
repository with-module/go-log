package log

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gitlab.com/with-junbach/go-modules/log/core"
	"testing"
	"time"
)

func TestCtxLog(t *testing.T) {
	err := LoadConfig(Config{
		Level:  "debug",
		Format: "json",
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, inst)
	defer Flush()

	CtxLog(context.Background(), "invalid-mode", "this log will print at debug level and use global inst due to empty ctx")
	log := Module("log-context")
	assert.IsType(t, new(core.BaseLogger), log)

	ctx := BindContext(context.Background(), "request_id", "some-id")
	assert.Implements(t, (*context.Context)(nil), ctx)
	log.CtxLog(ctx, "info", "print info message with request-id and latency duration", "latency", time.Millisecond*200)
}
