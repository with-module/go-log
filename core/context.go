package core

import "context"

type contextKey string

const contextKeyLogger contextKey = "CTX_KEY_LOGGER"

func (b *BaseLogger) Ctx(ctx context.Context) *BaseLogger {
	ctxLogger, ok := ctx.Value(contextKeyLogger).(*BaseLogger)
	if ok {
		return ctxLogger
	}
	return b
}

func (b *BaseLogger) CtxLog(ctx context.Context, mode string, msg string, args ...any) {
	b.Ctx(ctx).PrintLog(mode, msg, args...)
}

func (b *BaseLogger) BindContext(ctx context.Context, fields ...any) context.Context {
	return context.WithValue(ctx, contextKeyLogger, &BaseLogger{b.Ctx(ctx).With(fields...)})
}
