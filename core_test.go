package log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStd(t *testing.T) {
	inst := Std()
	inst.Info().Msg("message call with std instance")

	Debug("[DEBUG] call log message")
	Info("[INFO] call log message")
	Warn(fmt.Sprintf("[WARN] call log message: %s", "acceptable error"))
	Error(fmt.Errorf("connection error"), "[ERROR] call log message")

	panicLogFn := func() {

		Panic(fmt.Errorf("oom error"), "[PANIC] panic error log")
	}
	assert.Panicsf(t, panicLogFn, "this function log got panic")
}
