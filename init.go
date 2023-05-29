package log

import (
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

type (
	Config struct {
		Module string `config:"module" yaml:"module" json:"module"` // define module name. eg. awsome-api-web-service
		Level  string `config:"level" yaml:"level" json:"level"`    // [debug, info, warn, error, fatal, panic]
		Output string `config:"output" yaml:"output" json:"output"` // [stdout, stderr]
	}

	WithOption = func(zl zerolog.Logger) zerolog.Logger

	Logger = zerolog.Logger
)

var std zerolog.Logger

const (
	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel
	WarnLevel  = zerolog.WarnLevel
	ErrorLevel = zerolog.ErrorLevel
	PanicLevel = zerolog.PanicLevel
)

func New(cfg Config, options ...WithOption) Logger {
	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		Error(err, "invalid log level, [debug] will be applied by default")
		level = DebugLevel
	}

	output := os.Stdout
	if strings.EqualFold(cfg.Output, "stderr") {
		output = os.Stderr
	}

	inst := zerolog.New(output).Level(level).With().
		Timestamp().
		Caller().
		Str("module", cfg.Module).
		Logger()
	for _, fn := range options {
		inst = fn(inst)
	}

	return inst
}

func LoadConfig(cfg Config, opts ...WithOption) {
	logger := New(cfg, opts...)
	std = logger
	std.Debug().Interface("data", cfg).Msg("log config has been loaded successfully")
}

func init() {
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.DurationFieldInteger = true
	std = zerolog.New(os.Stdout).Level(DebugLevel).With().
		Timestamp().
		Caller().
		Str("module", "log-service").
		Logger()
	Debug("initiate logger with default settings")
}
