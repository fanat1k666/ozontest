package pkg

import (
	"context"
	"fmt"
)

type Logger interface {
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

var loggerContextKey = struct{}{}

func GetFromContext(ctx context.Context) (Logger, error) {
	l := ctx.Value(loggerContextKey)
	if l == nil {
		return nil, fmt.Errorf("logger isn't set")
	}

	tl, ok := l.(Logger)
	if !ok {
		return nil, fmt.Errorf("logger have wrong type")
	}

	return tl, nil
}

func AddToContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, l)
}
