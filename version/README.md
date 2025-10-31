# Version

A simple Go package for embedding and logging build and version information at compile time using slog.

## Example Usage

```go
package main

import (
    "log/slog"
    "os"

    "github.com/veerendra2/gopackages/version"
)

func main() {
    slog.Info("Version information", version.Info()...)
    slog.Info("Build context", version.BuildContext()...)
}
```

## Build with ldflags

Make sure to add flags while building the app

```bash
# Example: https://github.com/veerendra2/go-project-template/blob/main/Taskfile.yml

BRANCH    ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILDTIME ?= $(shell date '+%Y-%m-%d@%H:%M:%S')
BUILDUSER ?= $(shell id -un)
REVISION  ?= $(shell git rev-parse HEAD)
VERSION   ?= $(shell git describe --tags)

go build -ldflags "\
  -X github.com/veerendra2/gopackages/version.Version=$(VERSION) \
  -X github.com/veerendra2/gopackages/version.Revision=$(REVISION) \
  -X github.com/veerendra2/gopackages/version.Branch=$(BRANCH) \
  -X github.com/veerendra2/gopackages/version.BuildUser=$(BUILDUSER) \
  -X github.com/veerendra2/gopackages/version.BuildDate=$(BUILDTIME)" \
  -o ./bin/myapp ./cmd/myapp
```
