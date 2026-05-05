// Package logging provides a configured slog.Logger and helpers for
// extracting it from request contexts.
package logging

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type ctxKey int

const loggerKey ctxKey = 0

// New returns a slog.Logger writing to stdout. JSON format in production,
// human-readable text format otherwise. The level filters incoming records.
func New(production bool, level slog.Level) *slog.Logger {
	return newWithWriter(os.Stdout, production, level)
}

func newWithWriter(w io.Writer, production bool, level slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{Level: level}
	var h slog.Handler
	if production {
		h = slog.NewJSONHandler(w, opts)
	} else {
		h = slog.NewTextHandler(w, opts)
	}
	return slog.New(h)
}

// IntoContext attaches a logger (typically request-scoped) to ctx.
func IntoContext(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

// FromContext returns the logger attached to ctx, or slog.Default if none.
func FromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return l
	}
	return slog.Default()
}
