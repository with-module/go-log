package logger

type (
	Config struct {
		// Log level: can be one of: debug, info, warn, error, panic
		Level string `config:"Level"`

		// Where log will be written: stdout, stderr, file path
		Output []string `config:"Output"`

		// Where to write error output: eg. stderr
		ErrOutput []string `config:"ErrOutput"`

		// Log output format: plaintext, json
		Format string `config:"format"`
	}
)
