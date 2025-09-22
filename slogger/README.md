# Slogger

A custom "Go package" to simplify logger initialization

## Example Usage

```go
package main

import (
	"log/slog"

	"github.com/alecthomas/kong"
	"github.com/veerendra2/gopackages/slogger"
)

var cli struct {
	Log slogger.Config `embed:"" prefix:"log." envprefix:"LOG_"`
}

func main() {
	kongCtx := kong.Parse(&cli)
	kongCtx.FatalIfErrorf(kongCtx.Error)

	slog.SetDefault(slogger.New(cli.Log))
	slog.Info("Info log")
	slog.Warn("Warning log")
	slog.Debug("Debug log")
	slog.Error("Error log")
}
```

```bash
go mod tidy

go run main.go -h
Usage: main [flags]

Flags:
  -h, --help                 Show context-sensitive help.
      --log.format="json"    Set the output format of the logs. Must be "console" or "json" ($LOG_FORMAT).
      --log.level=INFO       Set the log level. Must be "DEBUG", "INFO", "WARN" or "ERROR" ($LOG_LEVEL).
      --log.add-source       Whether to add source file and line number to log records ($LOG_ADD_SOURCE).

go run main.go --log.level=debug
{"time":"2025-09-22T12:57:02+02:00","level":"INFO","source":{"function":"main.main","file":"/Users/veerendra.kakumanu/projects/gopackages/main.go","line":19},"msg":"Info log"}
{"time":"2025-09-22T12:57:02+02:00","level":"WARN","source":{"function":"main.main","file":"/Users/veerendra.kakumanu/projects/gopackages/main.go","line":20},"msg":"Warning log"}
{"time":"2025-09-22T12:57:02+02:00","level":"DEBUG","source":{"function":"main.main","file":"/Users/veerendra.kakumanu/projects/gopackages/main.go","line":21},"msg":"Debug log"}
{"time":"2025-09-22T12:57:02+02:00","level":"ERROR","source":{"function":"main.main","file":"/Users/veerendra.kakumanu/projects/gopackages/main.go","line":22},"msg":"Error log"}
```
