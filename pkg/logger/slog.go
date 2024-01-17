package logger

import (
	"log/slog"
	"os"
	"time"
)

func NewSlog(level string) *slog.Logger {
	opts := options(level)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &opts))

	return logger
}

func toLogLevel(lvl string) slog.Level {
	switch lvl {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func options(level string) slog.HandlerOptions {
	var opts slog.HandlerOptions

	opts.Level = toLogLevel(level)
	opts.ReplaceAttr = replaceAttr()

	return opts
}

func replaceAttr() func(groups []string, a slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			timeFormat := a.Value.Time().Format(time.RFC3339)
			a.Value = slog.StringValue(timeFormat)
		}

		return a
	}
}
