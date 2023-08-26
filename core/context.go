package core

import "context"

type contextKey string

const contextKeyLogger contextKey = "CTX_KEY_LOGGER"

func (c *Logger) Ctx(ctx context.Context) *Logger {
	ctxLogger, ok := ctx.Value(contextKeyLogger).(*Logger)
	if ok {
		return ctxLogger
	}
	return c
}

func (c *Logger) CtxLog(ctx context.Context, mode string, msg string, args ...any) {
	c.Ctx(ctx).PrintLog(mode, msg, args...)
}

func (c *Logger) BindContext(ctx context.Context, fields ...any) context.Context {
	return context.WithValue(ctx, contextKeyLogger, &Logger{c.Ctx(ctx).With(fields...)})
}
