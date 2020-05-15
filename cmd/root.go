package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/thedevsaddam/docgen/collection"
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

	extraCSS string

	GitCommit = "unknown"
	Version   = "unknown"
	BuildDate = "unknown"

	cmd = &cobra.Command{
		Use:   "docgen",
		Short: "Generate documentation from Postman JSON collection",
		Long:  logo,
	}
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

	// register commands
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(serveLive)
	cmd.AddCommand(buildOutput)
}

// Execute the root command
func Execute() error {
	return cmd.Execute()
}

func handleErr(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func readJSONtoHTML(str string) *bytes.Buffer {
	var rt collection.Documentation
	f, err := os.Open(str)
	if err != nil {
		log.Fatal("opening file", err.Error())
	}

	// open and build collection
	if err = rt.Open(f); err != nil {
		log.Fatal("parsing json file", err.Error())
	}

	// populate envCollection with collection variables
	if len(rt.Variables) > 0 {
		envCollection.SetCollectionVariables(rt.Variables)
	}

	// override collection variables by env variables
	if env != "" {
		if _, err := os.Stat(env); os.IsNotExist(err) {
			handleErr("Invalid environment file path", err)
		}
		f, err := os.Open(env)
		if err != nil {
			handleErr("Unable to open file", err)
		}
		if err := envCollection.Open(f); err != nil {
			handleErr("Unable to parse environment file", err)
		}
	}

	tm := template.New("main")
	tm.Delims("@{{", "}}@")
	tm.Funcs(template.FuncMap{
		"html":            html,
		"css":             css,
		"js":              js,
		"snake":           snake,
		"addOne":          addOne,
		"color":           color,
		"trimQueryParams": trimQueryParams,
		"date_time":       dateTime,
		"markdown":        markdown,
		"e":               e,
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

	// populate envCollection with collection variables
	if len(rt.Variables) > 0 {
		envCollection.SetCollectionVariables(rt.Variables)
	}

	// override collection variables by env variables
	if env != "" {
		if _, err := os.Stat(env); os.IsNotExist(err) {
			handleErr("Invalid environment file path", err)
		}
		f, err := os.Open(env)
		if err != nil {
			handleErr("Unable to open file", err)
		}
		if err := envCollection.Open(f); err != nil {
			handleErr("Unable to parse environment file", err)
		}
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
		"roman":           roman,
		"date_time":       dateTime,
		"trimQueryParams": trimQueryParams,
		"e":               e,
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
	var rt collection.Documentation
	f, err := os.Open(str)
	if err != nil {
		log.Fatal("opening file", err.Error())
	}
	if err = rt.Open(f); err != nil {
		log.Fatal("parsing json file", err.Error())
	}

	// populate envCollection with collection variables
	if len(rt.Variables) > 0 {
		envCollection.SetCollectionVariables(rt.Variables)
	}

	// override collection variables by env variables
	if env != "" {
		if _, err := os.Stat(env); os.IsNotExist(err) {
			handleErr("Invalid environment file path", err)
		}
		f, err := os.Open(env)
		if err != nil {
			handleErr("Unable to open file", err)
		}
		if err := envCollection.Open(f); err != nil {
			handleErr("Unable to parse environment file", err)
		}
	}

	tm := template.New("main")
	tm.Delims("@{{", "}}@")
	tm.Funcs(template.FuncMap{
		"html":            html,
		"css":             css,
		"js":              js,
		"snake":           snake,
		"addOne":          addOne,
		"color":           color,
		"trimQueryParams": trimQueryParams,
		"date_time":       dateTime,
		"markdown":        markdown,
		"e":               e,
	})

	t, err := tm.Parse(assets.MarkdownHTML)
	if err != nil {
		log.Fatal(err)
	}

	buf := readJSONtoMarkdown(file)
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
		Data         collection.Documentation
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
