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
		const logo = `
  _____                _____
 |  __ \              / ____|
 | |  | | ___   ___  | |  __  ___ _ __
 | |  | |/ _ \ / __| | | |_ |/ _ \ '_ \
 | |__| | (_) | (__  | |__| |  __/ | | |
 |_____/ \___/ \___|  \_____|\___|_| |_|
                                                                                
`
		fmt.Println(logo)
		fmt.Println("Docgen version: v2.1")
		fmt.Println("Support postman collection version > 2.1")
		fmt.Println("For more info visit: https://github.com/thedevsaddam/docgen")
	},
}
