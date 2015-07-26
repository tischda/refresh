#!/bin/sh
# --------------------------------------------------------------------------
# Build script
# --------------------------------------------------------------------------

go clean

# doc says not to use -s to remove the debug info read by gdb
# but https://github.com/golang/go/issues/6245 says it's fixed
go build -race -ldflags "-X main.version $(git describe --tags)" -s
