package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

const (
	JsonFormat      = "json"
	PlainTextFormat = "console"
)

func InitLogger(conf Config, opts ...zap.Option) (*BaseLogger, error) {
	level, err := zapcore.ParseLevel(conf.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}

	format := PlainTextFormat
	encoderLevel := zapcore.CapitalColorLevelEncoder
	if strings.EqualFold(conf.Format, JsonFormat) {
		format = JsonFormat
		encoderLevel = zapcore.CapitalLevelEncoder
	}

	encoder := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "module",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encoderLevel,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapConf := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          format,
		EncoderConfig:     encoder,
		OutputPaths:       conf.Writer.Output,
		ErrorOutputPaths:  conf.Writer.Error,
		InitialFields:     nil,
	}

	core, err := zapConf.Build(opts...)
	if err != nil {
		return nil, err
	}
	return &BaseLogger{core.Sugar()}, nil
}
