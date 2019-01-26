package main

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	assets Assets

	cmd = &cobra.Command{
		Use:   "docgen",
		Short: "Generate documentation from Postman JSON collection",
		Long:  `Generate documentation from Postman JSON collection`,
	}
)

const (
	defaultCollection = "Default"
)

func init() {
	// init assets
	assets = Assets{
		IndexHTML: getData("assets/index.html"),

		BootstrapJS:  getData("assets/bootstrap.min.js"),
		BootstrapCSS: getData("assets/bootstrap.min.css"),

		JqueryJS:  getData("assets/jquery.min.js"),
		ScriptsJS: getData("assets/scripts.js"),
		StylesCSS: getData("assets/styles.css"),

		IndexMarkdown: getData("assets/index.md"),
	}
	// register commands
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(serveLive)
	cmd.AddCommand(buildOutput)
}

func readJSONtoHTML(str string) *bytes.Buffer {
	var rt Root
	f, err := os.Open(str)
	if err != nil {
		log.Fatal("opening file", err.Error())
	}
	if err = rt.Open(f); err != nil {
		log.Fatal("parsing json file", err.Error())
	}

	tm := template.New("main")
	tm.Delims("@{{", "}}@")
	tm.Funcs(template.FuncMap{
		"html":            html,
		"css":             css,
		"js":              js,
		"snake":           snake,
		"color":           color,
		"trimQueryParams": trimQueryParams,
		"markdown":        markdown,
	})
	t, err := tm.Parse(assets.IndexHTML)
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Assets Assets
		Data   Root
	}{
		Assets: assets,
		Data:   rt,
	}
	buf := new(bytes.Buffer)
	// defer buf.Reset()
	if err := t.Execute(buf, data); err != nil {
		log.Fatal(err)
	}
	return buf
}
func readJSONtoMarkdown(str string) *bytes.Buffer {
	var rt Root
	f, err := os.Open(str)
	if err != nil {
		log.Fatal("opening file", err.Error())
	}
	if err = rt.Open(f); err != nil {
		log.Fatal("parsing json file", err.Error())
	}

	tm := template.New("main")
	tm.Delims("@{{", "}}@")
	tm.Funcs(template.FuncMap{
		"snake":           snake,
		"addOne":          addOne,
		"trimQueryParams": trimQueryParams,
	})
	t, err := tm.Parse(assets.IndexMarkdown)
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Data Root
	}{
		Data: rt,
	}
	buf := new(bytes.Buffer)
	// defer buf.Reset()
	if err := t.Execute(buf, data); err != nil {
		log.Fatal(err)
	}
	return buf
}

func main() {
	cmd.Execute()
}
