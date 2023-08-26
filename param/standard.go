package param

import "go.uber.org/zap"

func Err(err error) any {
	return zap.Error(err)
}

func Str(key, val string) any {
	return zap.String(key, val)
}

func Obj(key string, val any) any {
	return zap.Any(key, val)
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Int[T Integer](key string, val T) any {
	return zap.Int64(key, int64(val))
}
