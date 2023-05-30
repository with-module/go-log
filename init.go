package log

import (
	"github.com/rs/zerolog"
	"os"
	"strings"
	"sync"
	"time"
)

type (
	Config struct {
		Module string `config:"module" yaml:"module" json:"module"` // define module name. eg. awesome-api-web-service
		Level  string `config:"level" yaml:"level" json:"level"`    // [debug, info, warn, error, fatal, panic]
		Output string `config:"output" yaml:"output" json:"output"` // [stdout, stderr]
		Caller struct {
			Enabled   bool `config:"Enabled"`
			SkipFrame int  `config:"SkipFrame"`
		} `config:"Caller"`
	}

	WithOption = func(Logger) Logger

	Logger  = zerolog.Logger
	Context = zerolog.Context
)

var (
	std   zerolog.Logger
	setup sync.Once
)

const (
	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel
	WarnLevel  = zerolog.WarnLevel
	ErrorLevel = zerolog.ErrorLevel
	FatalLevel = zerolog.FatalLevel
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

	inst := zerolog.New(output).Level(level)
	for _, fn := range options {
		inst = fn(inst)
	}

	ctx := inst.With().
		Timestamp().
		Str("module", cfg.Module)
	if cfg.Caller.Enabled {
		ctx = ctx.CallerWithSkipFrameCount(cfg.Caller.SkipFrame)
	}
	return ctx.Logger()
}

func LoadConfig(cfg Config, opts ...WithOption) {
	setup.Do(func() {
		logger := New(cfg, opts...)
		std = logger
		zerolog.DefaultContextLogger = &std
		std.Debug().Interface("data", cfg).Msg("log config has been loaded successfully")
	})
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
