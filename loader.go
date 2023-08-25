package log

import (
	"gitlab.com/with-junbach/go-modules/log/param"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
)

const (
	JsonFormat      = "json"
	PlainTextFormat = "console"

	WriterConsole      = "stdout"
	WriterConsoleError = "stderr"

	DefaultLogLevel = "info"
)

func LoadConfig(conf Config, opts ...zap.Option) {
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
		log.Panicf("failed to initialize logger inst: %s", err)
	}

	inst = &Logger{core.Sugar()}
	inst.Debugw("logger instance has been initialized successfully", param.Obj("config", conf))
}

func init() {
	LoadConfig(Config{
		Level:  DefaultLogLevel,
		Format: PlainTextFormat,
		Writer: WriterConfig{
			Output: []string{WriterConsole},
			Error:  []string{WriterConsoleError},
		},
	})
}
