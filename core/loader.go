package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

const (
	FormatJson      = "json"
	FormatPlainText = "console"

	WriterConsole      = "stdout"
	WriterConsoleError = "stderr"
)

type (
	Config struct {
		Level  string       `json:"level" yaml:"level" config:"level"`
		Format string       `json:"format" yaml:"format" config:"format"`
		Writer WriterConfig `json:"writer" yaml:"writer" config:"writer"`
	}

	WriterConfig struct {
		Output []string `json:"output" yaml:"output" config:"output"`
		Error  []string `json:"error" yaml:"error" config:"error"`
	}
)

func NewLogger(conf Config, opts ...zap.Option) (*Logger, error) {
	level, err := zapcore.ParseLevel(conf.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}

	format := FormatPlainText
	encoderLevel := zapcore.CapitalColorLevelEncoder
	if strings.EqualFold(conf.Format, FormatJson) {
		format = FormatJson
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

	builder := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          format,
		EncoderConfig:     encoder,
		OutputPaths:       []string{WriterConsole},
		ErrorOutputPaths:  []string{WriterConsoleError},
		InitialFields:     nil,
	}

	if writers := conf.Writer.Output; len(writers) > 0 {
		builder.OutputPaths = writers
	}
	if writers := conf.Writer.Error; len(writers) > 0 {
		builder.ErrorOutputPaths = writers
	}

	core, err := builder.Build(opts...)
	if err != nil {
		return nil, err
	}
	return &Logger{core.Sugar()}, nil
}
