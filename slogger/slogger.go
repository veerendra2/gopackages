package slogger

import (
	"log/slog"
	"os"
)

// These tags are used by kong CLI argument parser.
type Config struct {
	Format    string     `env:"FORMAT" enum:"console,json" default:"json" help:"Set the output format of the logs. Must be \"console\" or \"json\"."`
	Level     slog.Level `env:"LEVEL" enum:"DEBUG,INFO,WARN,ERROR" default:"INFO" help:"Set the log level. Must be \"DEBUG\", \"INFO\", \"WARN\" or \"ERROR\"."`
	AddSource bool       `env:"ADD_SOURCE" default:"true" help:"Whether to add source file and line number to log records."`
}

func New(config Config) *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: config.AddSource,
		Level:     config.Level,
	}

	var handler slog.Handler
	if config.Format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
