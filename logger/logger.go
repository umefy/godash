package logger

import (
	"context"
	"io"
	"log/slog"
)

type LoggerOps struct {
	*slog.HandlerOptions
	Writer          io.Writer
	JSON            bool
	Level           slog.Level
	AddSource       bool
	SourceFieldName string
	CallerSkip      int // Adjust this based on your debugCaller() output
}

type Logger struct {
	*slog.Logger
	handler *ctxLoggerHandler
	opts    *LoggerOps
}

func NewLoggerOps(JSON bool, Writer io.Writer, level slog.Level, addSource bool, sourceFieldName string, callerSkip int) *LoggerOps {
	return &LoggerOps{
		HandlerOptions: &slog.HandlerOptions{
			Level: level,
		},
		Writer:          Writer,
		JSON:            JSON,
		Level:           level,
		AddSource:       addSource,
		SourceFieldName: sourceFieldName,
		CallerSkip:      callerSkip,
	}
}

func New(opts *LoggerOps, wrapHandler func(handler slog.Handler) slog.Handler) *Logger {
	var handler slog.Handler

	if opts.JSON {
		handler = slog.NewJSONHandler(opts.Writer, opts.HandlerOptions)
	} else {
		handler = slog.NewTextHandler(opts.Writer, opts.HandlerOptions)
	}

	if wrapHandler != nil {
		handler = wrapHandler(handler)
	}

	ctxHandler := &ctxLoggerHandler{Handler: handler, AddSource: opts.AddSource, SourceFieldName: opts.SourceFieldName, CallerSkip: opts.CallerSkip}

	return &Logger{
		Logger:  slog.New(ctxHandler),
		handler: ctxHandler,
		opts:    opts,
	}
}

func (l *Logger) GetOpts() *LoggerOps {
	return l.opts
}

func (l *Logger) GetHandler() ctxLoggerHandler {
	return *l.handler
}

func (l *Logger) WithValue(parent context.Context, attrs ...slog.Attr) context.Context {
	return l.handler.WithValue(parent, attrs...)
}
