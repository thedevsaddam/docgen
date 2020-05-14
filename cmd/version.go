package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Provide version information",
	Long:  `Provide version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(logo)
		fmt.Println("Version:", Version)
		fmt.Println("Git commit:", GitCommit)
		fmt.Println("Build date:", BuildDate)
		fmt.Println("Support postman collection version > 2.0")
	},
}
