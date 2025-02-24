# Logger

`logger` is a package that provides a logger that compatible with `slog` for the application.

## Usage

```go
func newLogger() *logger.Logger {
	loggerOpts := logger.NewLoggerOps(true, os.Stdout, slog.LevelDebug, true, "source", 4)
	logger := logger.New(loggerOpts, func(handler slog.Handler) slog.Handler {
		return handler.WithAttrs([]slog.Attr{
			slog.Int("pid", os.Getpid()),
		})
	})

	return logger
}
```
