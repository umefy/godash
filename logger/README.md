# logger

A powerful structured logging package for Go applications with enhanced context support and full compatibility with Go's `slog` package. Provides flexible logging configuration, context-aware logging, and middleware support for production applications.

## Features

- **slog Compatible**: Full compatibility with Go's structured logging standard
- **Context Support**: Enhanced context-aware logging with custom attributes
- **Flexible Configuration**: JSON and text output formats
- **Source Location**: Configurable source file and line tracking
- **Middleware Support**: Custom handler wrapping for global attributes
- **Performance Optimized**: Efficient logging with minimal overhead

## Installation

```bash
go get github.com/umefy/godash/logger
```

## Quick Start

```go
package main

import (
    "context"
    "os"
    "log/slog"
    "github.com/umefy/godash/logger"
)

func main() {
    // Create logger options
    opts := logger.NewLoggerOps(
        true,           // JSON output
        os.Stdout,      // Writer
        slog.LevelInfo, // Log level
        true,           // Add source location
        "source",       // Source field name
        4,              // Caller skip
    )

    // Create logger with middleware
    log := logger.New(opts, func(handler slog.Handler) slog.Handler {
        return handler.WithAttrs([]slog.Attr{
            slog.Int("pid", os.Getpid()),
            slog.String("service", "myapp"),
        })
    })

    // Basic logging
    log.Info("Application started", slog.String("version", "1.0.0"))

    // Context-aware logging
    ctx := log.WithValue(context.Background(), slog.String("request_id", "123"))
    log.InfoContext(ctx, "Request processed", slog.String("user", "alice"))
}
```

## API Reference

### LoggerOps Configuration

```go
type LoggerOps struct {
    *slog.HandlerOptions
    Writer          io.Writer
    JSON            bool
    Level           slog.Level
    AddSource       bool
    SourceFieldName string
    CallerSkip      int
}
```

#### NewLoggerOps

```go
func NewLoggerOps(JSON bool, Writer io.Writer, level slog.Level, addSource bool, sourceFieldName string, callerSkip int) *LoggerOps
```

Creates logger options with the specified configuration.

**Parameters:**

- `JSON`: Enable JSON output format (false for text)
- `Writer`: Output destination (e.g., `os.Stdout`, `os.Stderr`)
- `level`: Minimum log level to output
- `addSource`: Include source file and line information
- `sourceFieldName`: Field name for source information
- `callerSkip`: Number of call frames to skip for source location

### Logger Creation

#### New

```go
func New(opts *LoggerOps, wrapHandler func(handler slog.Handler) slog.Handler) *Logger
```

Creates a new logger instance with optional middleware.

**Example:**

```go
// Simple logger
log := logger.New(opts, nil)

// Logger with middleware
log := logger.New(opts, func(handler slog.Handler) slog.Handler {
    return handler.WithAttrs([]slog.Attr{
        slog.String("environment", "production"),
        slog.String("service", "api"),
    })
})
```

### Context Support

#### WithValue

```go
func (l *Logger) WithValue(parent context.Context, attrs ...slog.Attr) context.Context
```

Creates a context with logging attributes that will be included in all log messages.

**Example:**

```go
// Create context with request attributes
ctx := log.WithValue(context.Background(),
    slog.String("request_id", "abc123"),
    slog.String("user_id", "user456"),
)

// All log messages in this context will include the attributes
log.InfoContext(ctx, "Processing request")
log.ErrorContext(ctx, "Request failed", slog.String("error", "timeout"))
```

### Logger Methods

The logger provides all standard slog methods:

```go
// Basic logging
log.Debug(msg, args...)
log.Info(msg, args...)
log.Warn(msg, args...)
log.Error(msg, args...)

// Context-aware logging
log.DebugContext(ctx, msg, args...)
log.InfoContext(ctx, msg, args...)
log.WarnContext(ctx, msg, args...)
log.ErrorContext(ctx, msg, args...)
```

## Configuration Examples

### Development Logger

```go
// Text format with debug level and source location
opts := logger.NewLoggerOps(
    false,          // Text output
    os.Stdout,      // Console output
    slog.LevelDebug, // Debug level
    true,           // Add source
    "source",       // Source field
    4,              // Caller skip
)
```

### Production Logger

```go
// JSON format with info level
opts := logger.NewLoggerOps(
    true,           // JSON output
    os.Stdout,      // Output destination
    slog.LevelInfo, // Info level
    false,          // No source location
    "",             // No source field
    0,              // No caller skip
)
```

### File Logger

```go
file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
opts := logger.NewLoggerOps(
    true,           // JSON output
    file,           // File output
    slog.LevelInfo, // Info level
    true,           // Add source
    "source",       // Source field
    4,              // Caller skip
)
```

## Advanced Usage

### Custom Middleware

```go
// Add request tracing
log := logger.New(opts, func(handler slog.Handler) slog.Handler {
    return &tracingHandler{Handler: handler}
})

type tracingHandler struct {
    slog.Handler
}

func (h *tracingHandler) Handle(ctx context.Context, r slog.Record) error {
    // Add trace ID if available
    if traceID := ctx.Value("trace_id"); traceID != nil {
        r.AddAttrs(slog.String("trace_id", traceID.(string)))
    }
    return h.Handler.Handle(ctx, r)
}
```

### Structured Logging

```go
// Log structured data
log.Info("User action",
    slog.String("action", "login"),
    slog.String("user_id", "123"),
    slog.String("ip", "192.168.1.1"),
    slog.Time("timestamp", time.Now()),
)

// Log errors with context
if err != nil {
    log.Error("Database operation failed",
        slog.String("operation", "insert"),
        slog.String("table", "users"),
        slog.String("error", err.Error()),
    )
}
```

### Performance Logging

```go
start := time.Now()
// ... perform operation
log.Info("Operation completed",
    slog.String("operation", "data_processing"),
    slog.Duration("duration", time.Since(start)),
    slog.Int("records_processed", count),
)
```

## Output Formats

### JSON Output

```json
{
  "time": "2024-01-15T10:30:00.123Z",
  "level": "INFO",
  "msg": "Request processed",
  "pid": 12345,
  "service": "api",
  "request_id": "abc123",
  "user": "alice",
  "source": "main.go:42"
}
```

### Text Output

```
2024-01-15T10:30:00.123Z INFO Request processed pid=12345 service=api request_id=abc123 user=alice source=main.go:42
```

## Best Practices

1. **Use Context for Request Scoping**: Always use context-aware logging for request-specific information
2. **Structured Data**: Use slog attributes for structured data instead of string concatenation
3. **Appropriate Log Levels**: Use debug for development, info for general flow, warn for issues, error for failures
4. **Performance**: Avoid expensive operations in log statements
5. **Sensitive Data**: Never log sensitive information like passwords or tokens

## Testing

```bash
go test ./logger/...
```

## Dependencies

- `log/slog` - Go's structured logging package
- `context` - Context support
- `io` - Writer interface

## License

MIT License - see [LICENSE](../LICENSE) for details.
