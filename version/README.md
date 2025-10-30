# Version

A simple Go package for embedding and logging build and version information at compile time.

## Example Usage

```go
package main

import (
	"log/slog"

	"github.com/veerendra2/gopackages/version"
)

func main() {
  logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

  logger.Info("Version information", version.Info())
  logger.Info("Build context", version.BuildContext())

  // time=2025-10-30T20:58:12.861Z level=INFO msg="Version information" release.version=v1.2.3 release.branch=main release.revision=abc123
  // time=2025-10-30T20:58:12.862Z level=INFO msg="Build context" build.go_version=go1.24.7 build.user=tester build.date=2025-10-30T12:00:00Z
}
```

Make sure add flags while build in app

```bash
BRANCH    ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILDTIME ?= $(shell date '+%Y-%m-%d@%H:%M:%S')
BUILDUSER ?= $(shell id -un)
REVISION  ?= $(shell git rev-parse HEAD)
VERSION   ?= $(shell git describe --tags)

go build -ldflags "\
  -X github.com/yourorg/yourrepo/pkg/version.Version={{.VERSION}} \
  -X github.com/yourorg/yourrepo/pkg/version.Revision={{.REVISION}} \
  -X github.com/yourorg/yourrepo/pkg/version.Branch={{.BRANCH}} \
  -X github.com/yourorg/yourrepo/pkg/version.BuildUser={{.BUILDUSER}} \
  -X github.com/yourorg/yourrepo/pkg/version.BuildDate={{.BUILDTIME}}" \
  -o ./bin/myapp ./cmd/myapp
```
