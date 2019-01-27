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

	lines := strings.Split(buf.String(), "\n")
	for i, l := range lines {
		if strings.HasPrefix(l, "<!---") && strings.HasSuffix(l, "-->") {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	// contents := strings.Join(lines, "\n")
	var contents string
	var ws int
	for _, l := range lines {
		if l == "" {
			ws++
		} else {
			ws = 0
		}
		if ws <= 2 {
			contents += "\n" + l
		}
	}

	if err := ioutil.WriteFile(out, []byte(contents), 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("Documentation successfully generated to %s", out)
}
