package param

import (
	"go.uber.org/zap"
	"time"
)

func Any(key string, val any) any {
	return zap.Any(key, val)
}
func Err(err error) any {
	return zap.Error(err)
}

func Str(key, val string) any {
	return zap.String(key, val)
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Int[T Integer](key string, val T) any {
	return zap.Int64(key, int64(val))
}

func Time(key string, val time.Time) any {
	return zap.Time(key, val)
}

func Dur(key string, val time.Duration) any {
	return zap.Duration(key, val)
}

func Stack(key string) any {
	return zap.Stack(key)
}
