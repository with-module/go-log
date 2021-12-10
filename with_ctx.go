package log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

type (

	// JsonCustomizer Helper to customize output log of object in ContextLogger's functions
	// This can be used to strip out unwanted/sensitive data from the log
	JsonCustomizer interface {
		Customize() interface{}
	}

	GetRequestIDFromContextFunction func(ctx context.Context) string
)

var requestIDFunction GetRequestIDFromContextFunction = nil

func ImportGetRequestIDFunction(fn GetRequestIDFromContextFunction) {
	if fn != nil {
		requestIDFunction = fn
	}
}

func KDebug(ctx context.Context, template string, args ...interface{}) {
	logWithContext(ctx, DebugLevel, nil, template, args)
}

func KDDebug(ctx context.Context, data interface{}, template string, args ...interface{}) {
	logWithContext(ctx, DebugLevel, data, template, args)
}

func KInfo(ctx context.Context, template string, args ...interface{}) {
	logWithContext(ctx, InfoLevel, nil, template, args)
}

func KDInfo(ctx context.Context, data interface{}, template string, args ...interface{}) {
	logWithContext(ctx, InfoLevel, data, template, args)
}

func KWarn(ctx context.Context, template string, args ...interface{}) {
	logWithContext(ctx, WarnLevel, nil, template, args)
}

func KDWarn(ctx context.Context, data interface{}, template string, args ...interface{}) {
	logWithContext(ctx, WarnLevel, data, template, args)
}

func KError(ctx context.Context, template string, args ...interface{}) {
	logWithContext(ctx, ErrorLevel, nil, template, args)
}

func KDError(ctx context.Context, data interface{}, template string, args ...interface{}) {
	logWithContext(ctx, ErrorLevel, data, template, args)
}

func KPanic(ctx context.Context, template string, args ...interface{}) {
	logWithContext(ctx, PanicLevel, nil, template, args)
}

func KDPanic(ctx context.Context, data interface{}, template string, args ...interface{}) {
	logWithContext(ctx, PanicLevel, data, template, args)
}

func logWithContext(ctx context.Context, level Level, data interface{}, template string, args []interface{}) {
	var message = template
	if len(args) > 0 {
		message = fmt.Sprintf(template, args...)
	}

	var requestID = ""
	if requestIDFunction != nil {
		requestID = requestIDFunction(ctx)
	}

	Print(level, message, zap.String("request_id", requestID), zap.Any("object_data", logWithCustomizer(data)))
}

// Transform output
func logWithCustomizer(v interface{}) interface{} {
	if data, ok := v.(JsonCustomizer); ok {
		return data.Customize()
	}
	return v
}
