package log

import "io"

func WithOutput(w io.Writer) WithOption {
	return func(l Logger) Logger {
		return l.Output(w)
	}
}
