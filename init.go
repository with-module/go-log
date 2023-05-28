package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
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

func New(cfg Config, options ...WithOption) (Logger, error) {
	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		return std, fmt.Errorf("failed to parse log level: %v", err)
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

	return inst, nil
}

func LoadConfig(cfg Config, opts ...WithOption) error {
	logger, err := New(cfg, opts...)
	if err != nil {
		return fmt.Errorf("failed to apply log config: %v", err)
	}
	std = logger
	std.Debug().Interface("data", cfg).Msg("log config has been loaded successfully")
	return nil
}

func init() {
	std = zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().
		Timestamp().
		Caller().
		Str("module", "log-service").
		Logger()
	Debug("initiate logger with default settings")
}
