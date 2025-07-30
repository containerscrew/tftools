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
	version       string
	goversion     = runtime.Version()
	goos          = runtime.GOOS
	goarch        = runtime.GOARCH
	showTags      bool
	showUnchanged bool
	compact       bool
	useMarkdown   bool
	useJson       bool
	metrics       bool
	prettyJSON    bool
	hideDetails   bool
)

func init() {
	summarizeCmd.Flags().BoolVarP(&showTags, "show-tags", "t", false, "Show resources with tag changes")
	summarizeCmd.Flags().BoolVarP(&showUnchanged, "show-unchanged", "u", false, "Show resources with no changes")
	summarizeCmd.Flags().BoolVarP(&compact, "compact", "c", false, "Use compact formatting")
	summarizeCmd.Flags().BoolVarP(&useMarkdown, "markdown", "m", false, "Use markdown formatting")
	summarizeCmd.Flags().BoolVarP(&useJson, "json", "j", false, "Use JSON output")
	summarizeCmd.Flags().BoolVarP(&metrics, "metrics", "s", false, "Output metrics")
	summarizeCmd.Flags().BoolVarP(&prettyJSON, "pretty-json", "p", false, "Pretty JSON output")
	summarizeCmd.Flags().BoolVarP(&hideDetails, "hide-details", "d", false, "Hide details of changes")
}

// summarizeCmd will parse the tf plan output json to scrape created|updated|deleted resources in a clear outout
var summarizeCmd = &cobra.Command{
	Use:   "summarize",
	Short: "Get a summary of terraform/terragrunt output",
	Long:  "Get a summary of terraform/terragrunt output plan (created|updated|destroyed...)",
	Run: func(cmd *cobra.Command, args []string) {
		if useMarkdown && useJson {
			fmt.Println("-m (Markdown output) and -j (JSON output) are mutually exclusive")
			os.Exit(1)
		}

		if metrics && !useJson {
			fmt.Println("Metric output can only be used with JSON output")
			os.Exit(1)
		}

		output, err := reader.Reader(os.Stdin)
		if err != nil {
			panic(err)
		}

		parser.Parser(output, showTags, showUnchanged, compact, useMarkdown, useJson, metrics, prettyJSON, hideDetails)
	},
}

// usageCmd will print some docs files in a pretty markdown render
var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "print usage",
	Long:  "print usage in a pretty markdown render using terminal. This require internet connection since it fetch usage.md from github url",
	Run: func(cmd *cobra.Command, args []string) {
		if len(version) == 0 {
			version = "main"
		}

		markdownrender.RenderUsage(version)
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
