/*
Copyright © 2022 Containerscrew
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"

	"github.com/spf13/cobra"
)

// printBanner will print an ascii banner with colors
func printBanner() {
	templ := `{{ .AnsiColor.BrightMagenta  }} {{ .Title "tftools" "" 2 }}{{ .AnsiColor.Default }}
   Author: github.com/containerscrew
   License: Apache License 2.0
   Now: {{ .Now "Monday, 2 Jan 2006" }}`
	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	fmt.Printf("\n\n")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tftools",
	Short: "Useful tools to work with terraform/terragrunt",
	Long:  `Useful tools to work with terraform/terragrunt in your daily life`,
	Run: func(cmd *cobra.Command, args []string) {
		printBanner()
		if err := cmd.Help(); err != nil {
			panic(err)
		}
	},
}

// Execute starts the cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("😢 %s\n", err.Error())
		os.Exit(1)
	}
}
