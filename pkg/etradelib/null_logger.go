package etradelib

import (
	"context"
	"golang.org/x/exp/slog"
)

func CreateNullLogger() *slog.Logger {
	return slog.New(&nullHandler{})
}

type nullHandler struct {
}

func (h *nullHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

func (h *nullHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *nullHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return &nullHandler{}
}

func (h *nullHandler) WithGroup(_ string) slog.Handler {
	return &nullHandler{}
}
