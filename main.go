//go:generate go run generate-asset.go

package main

import (
	"github.com/thedevsaddam/docgen/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
