package slogger

import (
	"log/slog"
	"os"
	"time"
)

// These tags are used by kong CLI argument parser.
type Config struct {
	Format    string     `env:"FORMAT" enum:"console,json" default:"console" help:"Set the output format of the logs. Must be \"console\" or \"json\"."`
	Level     slog.Level `env:"LEVEL" enum:"DEBUG,INFO,WARN,ERROR" default:"INFO" help:"Set the log level. Must be \"DEBUG\", \"INFO\", \"WARN\" or \"ERROR\"."`
	AddSource bool       `env:"ADD_SOURCE" default:"false" help:"Whether to add source file and line number to log records."`
}

func New(config Config) *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: config.AddSource,
		Level:     config.Level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "time" {
				// Format time as RFC3339
				if t, ok := a.Value.Any().(time.Time); ok {
					return slog.Attr{
						Key:   a.Key,
						Value: slog.StringValue(t.Format(time.RFC3339)),
					}
				}
			}
			return a
		},
	}

	var handler slog.Handler

	handler = slog.NewTextHandler(os.Stdout, opts)
	if config.Format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
