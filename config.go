package log

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
