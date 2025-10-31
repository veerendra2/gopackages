package main

import (
	"fmt"
	"log/slog"
	"os"
)

var (
	GoVersion = "1.23"
	BuildUser = "root"
	BuildDate = "1.2.2022"
)

func BuildContext() []any {
	return []any{
		"go_version", GoVersion,
		"user", BuildUser,
		"date", BuildDate,
	}
}

func main() {
	// Set up a logger with JSON output for better visibility
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	fmt.Println(BuildContext()...)
	// Test the BuildContext function
	slog.Info("Build context", BuildContext()...)

	// // You can also test with additional context
	// slog.Info("Application starting",
	// 	append([]any{"app", "test-app"}, BuildContext()...)...)
}
