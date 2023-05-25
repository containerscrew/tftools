package cmd

import (
	"fmt"
	"os"
	"runtime"

	markdownrender "github.com/containerscrew/tftools/internal/utils/markdown_render"

	"github.com/containerscrew/tftools/internal/parser"
	"github.com/containerscrew/tftools/internal/reader"
	"github.com/spf13/cobra"
)

var (
	version   string
	goversion = runtime.Version()
	goos      = runtime.GOOS
	goarch    = runtime.GOARCH
)

// summarizeCmd will parse the tf plan output json to scrape created|updated|deleted resources in a clear outout
var summarizeCmd = &cobra.Command{
	Use:   "summarize",
	Short: "Get a summary of terraform/terragrunt output",
	Long:  "Get a summary of terraform/terragrunt output plan (created|updated|destroyed...)",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := reader.Reader(os.Stdin)
		if err != nil {
			panic(err)
		}

		parser.Parser(output)
	},
}

// usageCmd will print some docs files in a pretty markdown render
var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "print usage",
	Long:  "print usage in a pretty markdown render using terminal. This require internet connection since it fetch usage.md from github url",
	Run: func(cmd *cobra.Command, args []string) {
		markdownrender.RenderUsage()
	},
}

// versionCmd will print the current installed version in your local
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "tftools current version",
	Long:  "Get the cli tftools version installed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tftools: %s with go version %s %s/%s", version, goversion, goos, goarch)
	},
}
