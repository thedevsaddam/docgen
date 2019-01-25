package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func buildAndGenerateMarkdownFile(cmd *cobra.Command, args []string) {
	if in == "" {
		log.Println("You must provide a input file name!")
		return
	}

	if out == "" {
		log.Println("You must provide a output file name!")
		return
	}
	if _, err := os.Stat(in); os.IsNotExist(err) {
		log.Println("Invalid file path!")
		return
	}
	buf := readJSONtoMarkdown(in)
	if !strings.HasSuffix(out, ".md") {
		out = out + ".md"
	}

	if err := ioutil.WriteFile(out, buf.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("Documentation successfully generated to %s", out)
}
