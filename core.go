package log

import "os"

func Debug(msg string) {
	std.Debug().Msg(msg)
}

func Info(msg string) {
	std.Info().Msg(msg)
}

func Warn(msg string) {
	std.Warn().Msg(msg)
}

func Error(err error, msg string) {
	std.Error().Err(err).Msg(msg)
}

func Panic(err error, msg string) {
	std.Panic().Err(err).Msg(msg)
}

var exitFn = os.Exit

func Fatal(err error, msg string) {
	defer exitFn(1)
	std.WithLevel(FatalLevel).Err(err).Msg(msg)
}

func Std() *Logger {
	return &std
}
