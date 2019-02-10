#!/bin/bash

#MAC OS
export GOARCH="386"
export GOOS="darwin"
export CGO_ENABLED=1
go build -o mac_386 -v

export GOARCH="amd64"
export GOOS="darwin"
export CGO_ENABLED=1
go build -o mac_amd64 -v

#LINUX
export GOARCH="amd64"
export GOOS="linux"
export CGO_ENABLED=0
go build -o linux_amd64 -v

export GOARCH="386"
export GOOS="linux"
export CGO_ENABLED=0
go build -o linux_386 -v

#FREEBSD
export GOARCH="amd64"
export GOOS="freebsd"
export CGO_ENABLED=0
go build -o freebsd_amd64 -v

export GOARCH="386"
export GOOS="freebsd"
export CGO_ENABLED=1
go build -o freebsd_386 -v

#OPENBSD
export GOARCH="amd64"
export GOOS="openbsd"
export CGO_ENABLED=0
go build -o freebsd_amd64 -v

export GOARCH="386"
export GOOS="openbsd"
export CGO_ENABLED=0
go build -o freebsd_386 -v

#WINDOWS
export GOARCH="386"
export GOOS="windows"
export CGO_ENABLED=0
go build -o windows_386.exe -v

export GOARCH="amd64"
export GOOS="windows"
export CGO_ENABLED=0
go build -o windows_amd64.exe -v
