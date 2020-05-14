#!/bin/bash

GIT_COMMIT=$(git rev-parse --short HEAD)
TAG=$(git describe --exact-match --abbrev=0 --tags ${COMMIT} 2> /dev/null || true)
DATE=$(date +'%Y-%m-%d')

echo "Building binaries"
echo Git commit: $GIT_COMMIT Version: $TAG Build date: $DATE

go generate

# MAC
export GOARCH="amd64"
export GOOS="darwin"
export CGO_ENABLED=1
go build -ldflags "-X github.com/thedevsaddam/docgen/cmd.GitCommit=$GIT_COMMIT -X github.com/thedevsaddam/docgen/cmd.Version=$TAG -X github.com/thedevsaddam/docgen/cmd.BuildDate=$DATE" -o mac_amd64 -v .

#LINUX
export GOARCH="amd64"
export GOOS="linux"
export CGO_ENABLED=0
go build -ldflags "-X github.com/thedevsaddam/docgen/cmd.GitCommit=$GIT_COMMIT -X github.com/thedevsaddam/docgen/cmd.Version=$TAG -X github.com/thedevsaddam/docgen/cmd.BuildDate=$DATE" -o linux_amd64 -v

export GOARCH="386"
export GOOS="linux"
export CGO_ENABLED=0
go build -ldflags "-X github.com/thedevsaddam/docgen/cmd.GitCommit=$GIT_COMMIT -X github.com/thedevsaddam/docgen/cmd.Version=$TAG -X github.com/thedevsaddam/docgen/cmd.BuildDate=$DATE" -o linux_386 -v

#WINDOWS
export GOARCH="386"
export GOOS="windows"
export CGO_ENABLED=0
go build -ldflags "-X github.com/thedevsaddam/docgen/cmd.GitCommit=$GIT_COMMIT -X github.com/thedevsaddam/docgen/cmd.Version=$TAG -X github.com/thedevsaddam/docgen/cmd.BuildDate=$DATE" -o windows_386.exe -v

export GOARCH="amd64"
export GOOS="windows"
export CGO_ENABLED=0
go build -ldflags "-X github.com/thedevsaddam/docgen/cmd.GitCommit=$GIT_COMMIT -X github.com/thedevsaddam/docgen/cmd.Version=$TAG -X github.com/thedevsaddam/docgen/cmd.BuildDate=$DATE" -o windows_amd64.exe -v

echo "Build complete"