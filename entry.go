package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type (
	Config struct {
		// Log level: can be one of: debug, info, warn, error, panic
		Level string `config:"Level"`

		// Where log will be written: stdout, stderr, file path
		Output []string `config:"Output"`

		// Where to write error output: eg. stderr
		ErrOutput []string `config:"ErrOutput"`

		// Log output format: console, json
		Format string `config:"Format"`

		// Module to name the log instance, set to hostname if empty
		Module string `config:"Module"`
	}

	Level = string
	Field = zap.Field
)

const (
	DebugLevel Level = "debug"
	InfoLevel  Level = "info"
	WarnLevel  Level = "warn"
	ErrorLevel Level = "error"
	PanicLevel Level = "panic"
)

var inst *zap.SugaredLogger

// InitLoggerInst Initialize global log instance using input Config
func InitLoggerInst(conf Config) error {
	logName := zap.String("module", conf.Module)
	if conf.Module == "" {
		hostname, err := os.Hostname()
		if err != nil {
			return fmt.Errorf("failed to retreive hostname: %v", err)
		}

		logName = zap.String("module", hostname)
	}

	level := zap.InfoLevel
	if err := level.Set(conf.Level); err != nil {
		return fmt.Errorf("invalid log level config: %v", err)
	}

	encoder := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapConfig := zap.Config{
		Level: zap.NewAtomicLevelAt(level),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		EncoderConfig:     encoder,
		Encoding:          conf.Format,
		OutputPaths:       conf.Output,
		ErrorOutputPaths:  conf.ErrOutput,
	}

	log, err := zapConfig.Build(zap.AddStacktrace(zap.ErrorLevel), zap.AddCallerSkip(1), zap.Fields(logName))
	if err != nil {
		return fmt.Errorf("failed to initialize log instance: %v", err)
	}
	inst = log.Sugar()
	return nil
}

func Debug(template string, args ...interface{}) {
	inst.Debugf(template, args...)
}

func Info(template string, args ...interface{}) {
	inst.Infof(template, args...)
}

func Warn(template string, args ...interface{}) {
	inst.Warnf(template, args...)
}

func Error(template string, args ...interface{}) {
	inst.Errorf(template, args...)
}

func Panic(template string, args ...interface{}) {
	inst.Panicf(template, args...)
}

func PrintMap(level string, message string, dataMap map[string]interface{}) {
	var keyAndValues = make([]interface{}, len(dataMap))
	for k, v := range dataMap {
		keyAndValues = append(keyAndValues, With(k, v))
	}
	Print(level, message, keyAndValues...)
}

func Print(level string, message string, keysAndValues ...interface{}) {
	var log func(msg string, fields ...interface{})
	switch level {
	case DebugLevel:
		log = inst.Debugw
	case InfoLevel:
		log = inst.Infow
	case WarnLevel:
		log = inst.Warnw
	case ErrorLevel:
		log = inst.Errorw
	case PanicLevel:
		log = inst.Panicw
	default:
		log = inst.Debugw
	}

	log(message, keysAndValues...)
}

func With(key string, value interface{}) Field {
	return zap.Any(key, value)
}

// Close function used to flush log buffer.
func Close() error {
	inst.Info("flush log buffer")
	return inst.Sync()
}
