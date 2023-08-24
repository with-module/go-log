package log

import "context"

type contextKey string

const contextKeyLogger contextKey = "CTX_KEY_LOGGER"

func getFromContext(ctx context.Context) *Logger {
	ctxLogger, ok := ctx.Value(contextKeyLogger).(*Logger)
	if ok {
		return ctxLogger
	}
	return nil
}

func BindContext(ctx context.Context, fields ...any) context.Context {
	return context.WithValue(ctx, contextKeyLogger, &Logger{Ctx(ctx).With(fields...)})
}

func Ctx(ctx context.Context) *Logger {
	return inst.Ctx(ctx)
}

func CtxLog(ctx context.Context, mode string, msg string, args ...any) {
	inst.CtxLog(ctx, mode, msg, args...)
}

func (c *Logger) Ctx(ctx context.Context) *Logger {
	ctxLogger := getFromContext(ctx)
	if ctxLogger == nil {
		return c
	}

	return ctxLogger
}

func (c *Logger) CtxLog(ctx context.Context, mode string, msg string, args ...any) {
	c.Ctx(ctx).PrintLog(mode, msg, args...)
}
