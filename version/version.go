package version

import (
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

// Info returns key-value pairs for application version information.
func Info() []any {
	return []any{
		"version", Version,
		"branch", Branch,
		"revision", Revision,
	}
}

// BuildContext returns key-value pairs for build environment details.
func BuildContext() []any {
	return []any{
		"go_version", GoVersion,
		"user", BuildUser,
		"date", BuildDate,
	}
}
