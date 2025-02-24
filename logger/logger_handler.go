package logger

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
)

type ctxKey struct{}

type ctxLoggerHandler struct {
	slog.Handler
	AddSource       bool
	SourceFieldName string
	CallerSkip      int // Adjust this based on your debugCaller() output
}

var _ slog.Handler = (*ctxLoggerHandler)(nil)

func (h *ctxLoggerHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(ctxKey{}).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	if h.AddSource {
		_, file, line, ok := runtime.Caller(h.CallerSkip)
		if ok {
			if h.SourceFieldName == "" {
				h.SourceFieldName = "source"
			}

			r.AddAttrs(slog.String(h.SourceFieldName, fmt.Sprintf("%s:%d", file, line)))
		}
	}
	return h.Handler.Handle(ctx, r)
}

func (h *ctxLoggerHandler) WithValue(parent context.Context, attrs ...slog.Attr) context.Context {
	if v, ok := parent.Value(ctxKey{}).([]slog.Attr); ok {
		v = append(v, attrs...)
		return context.WithValue(parent, ctxKey{}, v)
	}

	return context.WithValue(parent, ctxKey{}, attrs)
}
