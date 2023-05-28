package log

import (
	"context"
	"fmt"
	"testing"
)

func TestBindContext(t *testing.T) {
	ctx := context.Background()
	bindingFields := map[string]any{
		"requestId":    "some-unique-id",
		"functionCode": 1001,
	}

	ctx = BindContext(ctx, bindingFields)
	KDebug(ctx, "[DEBUG] message log by context processing")
	KInfo(ctx, "[INFO] message log by context processing")
	KWarn(ctx, "[WARN] message log by context processing")

	KError(ctx, fmt.Errorf("handler error"), "[ERROR] message log by context processing")
	Ctx(ctx).Warn().Str("statusMessage", "failed due to invalid input").Msg("additional custom message")

}
