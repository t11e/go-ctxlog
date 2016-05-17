package ctxlog

import (
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

type contextKey int

const (
	loggerKey contextKey = iota
)

func FromContext(ctx context.Context) *logging.Logger {
	log, ok := ctx.Value(loggerKey).(*logging.Logger)
	if !ok {
		return Discard.Logger("context")
	}
	return log
}

func NewContext(ctx context.Context, log *logging.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}
