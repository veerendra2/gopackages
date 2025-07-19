package slogger

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name           string
		config         Config
		wantSource     bool
		wantLevel      slog.Level
		wantJSONFormat bool
		logFunc        func(*slog.Logger, string) // Add this to use different log levels
	}{
		{
			name: "JSON format with source",
			config: Config{
				Format:    "json",
				Level:     slog.LevelInfo,
				AddSource: true,
			},
			wantSource:     true,
			wantLevel:      slog.LevelInfo,
			wantJSONFormat: true,
			logFunc:        func(l *slog.Logger, msg string) { l.Info(msg) },
		},
		{
			name: "JSON format without source",
			config: Config{
				Format:    "json",
				Level:     slog.LevelInfo,
				AddSource: false,
			},
			wantSource:     false,
			wantLevel:      slog.LevelInfo,
			wantJSONFormat: true,
			logFunc:        func(l *slog.Logger, msg string) { l.Info(msg) },
		},
		{
			name: "Console format with source",
			config: Config{
				Format:    "console",
				Level:     slog.LevelInfo,
				AddSource: true,
			},
			wantSource:     true,
			wantLevel:      slog.LevelInfo,
			wantJSONFormat: false,
			logFunc:        func(l *slog.Logger, msg string) { l.Info(msg) },
		},
		{
			name: "JSON format with DEBUG level",
			config: Config{
				Format:    "json",
				Level:     slog.LevelDebug,
				AddSource: true,
			},
			wantSource:     true,
			wantLevel:      slog.LevelDebug,
			wantJSONFormat: true,
			logFunc:        func(l *slog.Logger, msg string) { l.Debug(msg) },
		},
		{
			name: "Console format with ERROR level",
			config: Config{
				Format:    "console",
				Level:     slog.LevelError,
				AddSource: true,
			},
			wantSource:     true,
			wantLevel:      slog.LevelError,
			wantJSONFormat: false,
			logFunc:        func(l *slog.Logger, msg string) { l.Error(msg) },
		},
		{
			name: "JSON format with WARN level",
			config: Config{
				Format:    "json",
				Level:     slog.LevelWarn,
				AddSource: false,
			},
			wantSource:     false,
			wantLevel:      slog.LevelWarn,
			wantJSONFormat: true,
			logFunc:        func(l *slog.Logger, msg string) { l.Warn(msg) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			// Use the New function to create the logger
			logger := New(tt.config, &buf)
			tt.logFunc(logger, "test message")

			output := buf.String()

			if !tt.wantJSONFormat {
				// For console format, check if output contains expected elements
				if tt.wantSource && !strings.Contains(output, "slogger_test.go") {
					t.Error("Expected source information in console output")
				}
				if !strings.Contains(output, strings.ToUpper(tt.wantLevel.String())) {
					t.Errorf("Expected level %s in console output", tt.wantLevel)
				}
				if !strings.Contains(output, "test message") {
					t.Error("Expected test message in console output")
				}
				return
			}

			// Parse JSON output
			var logEntry map[string]interface{}
			if err := json.Unmarshal([]byte(output), &logEntry); err != nil {
				t.Fatalf("Failed to parse JSON output: %v", err)
			}

			// Check if source is present when expected
			_, hasSource := logEntry["source"]
			if hasSource != tt.wantSource {
				t.Errorf("Source presence = %v, want %v", hasSource, tt.wantSource)
			}

			// Verify log level
			if level := logEntry["level"]; level != tt.wantLevel.String() {
				t.Errorf("Log level = %v, want %v", level, tt.wantLevel)
			}

			// Verify message
			if msg := logEntry["msg"]; msg != "test message" {
				t.Errorf("Message = %v, want 'test message'", msg)
			}
		})
	}
}
