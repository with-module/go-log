package log

type (
	Config struct {
		Level       string   `json:"level" yaml:"level" config:"level"`
		Format      string   `json:"format" yaml:"format" config:"format"`
		Output      []string `json:"output" yaml:"output" config:"output"`
		ErrorOutput []string `json:"errorOutput" yaml:"errorOutput" config:"errorOutput"`
	}
)

const (
	JsonFormat      = "json"
	PlainTextFormat = "console"

	ConsoleOutput      = "stdout"
	ConsoleOutputError = "stderr"
)
