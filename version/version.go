package version

import (
	"log/slog"
	"runtime"
)

var (
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	GoVersion = runtime.Version()
)

// Info returns slog attributes for application version information.
func Info() []any {
	return []any{
		slog.String("version", Version),
		slog.String("branch", Branch),
		slog.String("revision", Revision),
	}
}

// BuildContext returns slog attributes for build environment details.
func BuildContext() []any {
	return []any{
		slog.String("go_version", GoVersion),
		slog.String("user", BuildUser),
		slog.String("date", BuildDate),
	}
}
