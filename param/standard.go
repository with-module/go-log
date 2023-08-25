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
