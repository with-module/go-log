package log

import (
	"context"
	"gitlab.com/with-junbach/go-modules/log/core"
)

func BindContext(ctx context.Context, fields ...any) context.Context {
	return inst.BindContext(ctx, fields...)
}

func Ctx(ctx context.Context) *core.BaseLogger {
	return inst.Ctx(ctx)
}

func CtxLog(ctx context.Context, mode string, msg string, args ...any) {
	Ctx(ctx).PrintLog(mode, msg, args...)
}
