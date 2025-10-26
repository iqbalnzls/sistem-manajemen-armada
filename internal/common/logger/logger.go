package logger

import (
	"context"

	"go.uber.org/zap"
)

type contextKey string

const (
	loggerKey contextKey = "logger"
)

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return logger
	}

	return zap.NewNop()
}
