package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Provide version information",
	Long:  `Provide version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Docgen version: v2.0")
		fmt.Println("Support postman collection version > 2.1")
	},
}
