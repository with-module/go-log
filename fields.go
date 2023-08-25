package log

import "go.uber.org/zap"

func AddError(err error) any {
	return zap.Error(err)
}

func AddString(key, val string) any {
	return zap.String(key, val)
}

func AddObject(key string, val any) any {
	return zap.Any(key, val)
}
