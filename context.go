package log

import (
	"context"
	"github.com/rs/zerolog"
)

func BindContext(ctx context.Context, bindingFields map[string]any) context.Context {
	return std.With().Fields(bindingFields).Logger().WithContext(ctx)
}

func Ctx(ctx context.Context) *Logger {
	return zerolog.Ctx(ctx)
}

func KDebug(ctx context.Context, msg string) {
	Ctx(ctx).Debug().Msg(msg)
}

func KInfo(ctx context.Context, msg string) {
	Ctx(ctx).Info().Msg(msg)
}

func KWarn(ctx context.Context, msg string) {
	Ctx(ctx).Warn().Msg(msg)
}

func KError(ctx context.Context, err error, msg string) {
	Ctx(ctx).Error().Err(err).Msg(msg)
}
