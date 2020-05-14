// +build generateasset

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = http.Dir("./assets/")
	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:     "./assets_bin/assets.go",
		PackageName:  "assets_bin",
		VariableName: "AssetFS",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
