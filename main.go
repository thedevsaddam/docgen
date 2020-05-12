//go:generate go run generate-asset.go

package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/thedevsaddam/docgen/collection"

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

	extraCSS = ""

	cmd = &cobra.Command{
		Use:   "docgen",
		Short: "Generate documentation from Postman JSON collection",
		Long:  logo,
	}
)

const (
	defaultCollection = "Default"
)

// Assets represents template assets
type Assets struct {
	BootstrapJS  string
	BootstrapCSS string
	IndexHTML    string
	JqueryJS     string
	ScriptsJS    string
	StylesCSS    string
	ExtraCSS     string

	IndexMarkdown        string
	MarkdownHTML         string
	GithubMarkdownMinCSS string
}

func init() {
	// init assets
	assets = Assets{
		IndexHTML: getData("index.html"),

		BootstrapJS:  getData("bootstrap.min.js"),
		BootstrapCSS: getData("bootstrap.min.css"),

		JqueryJS:  getData("jquery.min.js"),
		ScriptsJS: getData("scripts.js"),
		StylesCSS: getData("styles.css"),

		IndexMarkdown:        getData("index.md"),
		MarkdownHTML:         getData("markdown.html"),
		GithubMarkdownMinCSS: getData("github-markdown.min.css"),
	}

	//Add flags
	cmd.PersistentFlags().StringVarP(&extraCSS, "css", "c", "", "Inject a css file")

	// register commands
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(serveLive)
	cmd.AddCommand(buildOutput)
}

func readJSONtoHTML(str string) *bytes.Buffer {
	var rt collection.Documentation
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

	if extraCSS != "" {
		ecFile, err := ioutil.ReadFile(extraCSS)
		if err != nil {
			log.Fatal("css file", err.Error())
		}
		assets.ExtraCSS = string(ecFile)
	}

	data := struct {
		Assets Assets
		Data   collection.Documentation
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
	var rt collection.Documentation
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
		Data collection.Documentation
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

	if extraCSS != "" {
		ecFile, err := ioutil.ReadFile(extraCSS)
		if err != nil {
			log.Fatal("css file", err.Error())
		}
		assets.ExtraCSS = string(ecFile)
	}

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
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
