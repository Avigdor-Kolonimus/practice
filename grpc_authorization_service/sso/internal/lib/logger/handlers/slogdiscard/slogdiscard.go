package slogdiscard

import (
	"context"

	"golang.org/x/exp/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

type DiscardHandler struct{}

func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	// Ignore the log record.
	return nil
}

func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	// Return the same handler; no attrs are stored.
	return h
}

func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	// Return the same handler; no group is stored.
	return h
}

func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	// Always false; log records are discarded.
	return false
}
