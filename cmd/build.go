package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thedevsaddam/docgen/collection"
)

var (
	in              string
	out             string
	env             string
	isMarkdown      bool
	includeVariable bool
	sorted          bool
	envCollection   = &collection.Environment{}
	buildOutput     = &cobra.Command{
		Use:   "build",
		Short: "Build html/markdown documentation from postman collection",
		Long:  `Build html/markdown documentation from postman collection`,
		Run:   buildAngGenerateFile,
	}
)

func init() {
	buildOutput.PersistentFlags().StringVarP(&in, "in", "i", "", "postman collection path")
	buildOutput.PersistentFlags().StringVarP(&out, "out", "o", "", "output file path")
	buildOutput.PersistentFlags().BoolVarP(&isMarkdown, "md", "m", false, "this flag will command to generate markdown")
	buildOutput.PersistentFlags().BoolVarP(&includeVariable, "var", "v", false, "this flag will include variables in template")
	buildOutput.PersistentFlags().BoolVarP(&sorted, "sort", "s", false, "this flag will sort the collection in ascending order")
	buildOutput.PersistentFlags().StringVarP(&extraCSS, "css", "c", "", "inject a css file")
	buildOutput.PersistentFlags().StringVarP(&env, "env", "e", "", "postman environment variable file path")
}

func buildAngGenerateFile(cmd *cobra.Command, args []string) {
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
	if isMarkdown {
		buildAndGenerateMarkdownFile(cmd, args)
		return
	}
	buildAndGenerateHTMLFile(cmd, args)
}
