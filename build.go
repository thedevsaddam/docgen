package main

import (
	"github.com/spf13/cobra"
)

var (
	in          string
	out         string
	isMarkdown  bool
	buildOutput = &cobra.Command{
		Use:   "build",
		Short: "Build html/markdown documentation from postman collection",
		Long:  `Build html/markdown documentation from postman collection`,
		Run:   buildAngGenerateFile,
	}
)

func init() {
	buildOutput.PersistentFlags().StringVarP(&in, "in", "i", "", "postman collection file relative path")
	buildOutput.PersistentFlags().StringVarP(&out, "out", "o", "", "output file relative path")
	buildOutput.PersistentFlags().BoolVarP(&isMarkdown, "md", "m", false, "this flag will command to generate markdown")
}

func buildAngGenerateFile(cmd *cobra.Command, args []string) {
	if isMarkdown {
		buildAndGenerateMarkdownFile(cmd, args)
		return
	}
	buildAndGenerateHTMLFile(cmd, args)
}
