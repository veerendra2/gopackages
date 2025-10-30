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

func Info() slog.Attr {
	return slog.Group("release",
		slog.String("version", Version),
		slog.String("branch", Branch),
		slog.String("revision", Revision),
	)
}

func BuildContext() slog.Attr {
	return slog.Group("build",
		slog.String("go_version", GoVersion),
		slog.String("user", BuildUser),
		slog.String("date", BuildDate),
	)
}
