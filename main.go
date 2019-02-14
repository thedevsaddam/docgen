package main

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

const logo = `
 _____                _____
|  __ \              / ____|
| |  | | ___   ___  | |  __  ___ _ __
| |  | |/ _ \ / __| | | |_ |/ _ \ '_ \
| |__| | (_) | (__  | |__| |  __/ | | |
|_____/ \___/ \___|  \_____|\___|_| |_|

Generate API documentation from Postman JSON collection
For more info visit: https://github.com/thedevsaddam/docgen
`

var (
	assets Assets

	githubLinkInc = make(map[string]int)

	cmd = &cobra.Command{
		Use:   "docgen",
		Short: "Generate documentation from Postman JSON collection",
		Long:  logo,
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

		IndexMarkdown:        getData("assets/index.md"),
		MarkdownHTML:         getData("assets/markdown.html"),
		GithubMarkdownMinCSS: getData("assets/github-markdown.min.css"),
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
		"date_time":       dateTime,
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
		"trim":            trim,
		"lower":           lower,
		"upper":           upper,
		"glink":           githubLink,
		"glinkInc":        githubLinkIncrementer,
		"merge":           merge,
		"date_time":       dateTime,
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
func readJSONtoMarkdownHTML(str string) *bytes.Buffer {
	// reset the link map
	githubLinkInc = make(map[string]int)
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
		"date_time":       dateTime,
		"markdown":        markdown,
	})
	t, err := tm.Parse(assets.MarkdownHTML)
	if err != nil {
		log.Fatal(err)
	}
	var buf *bytes.Buffer
	buf = readJSONtoMarkdown(file)
	mdHTML := markdown(buf.String())

	data := struct {
		Assets       Assets
		Data         Root
		MarkdownHTML string
	}{
		Assets:       assets,
		Data:         rt,
		MarkdownHTML: mdHTML,
	}
	buf.Reset()
	if err := t.Execute(buf, data); err != nil {
		log.Fatal(err)
	}
	return buf
}
func main() {
	cmd.Execute()
}
