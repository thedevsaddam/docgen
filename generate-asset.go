// +build generateasset

package main

import (
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
)

func main() {
	var fs http.FileSystem = http.Dir("./assets/")
	err := vfsgen.Generate(fs, vfsgen.Options{
		VariableName: "AssetFS",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
