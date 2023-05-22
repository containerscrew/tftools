package cmd

import (
	"os"

	markdownrender "github.com/containerscrew/tftools/internal/utils/markdown_render"

	"github.com/containerscrew/tftools/internal/parser"
	"github.com/containerscrew/tftools/internal/reader"
	"github.com/spf13/cobra"
)

//func Test() *cobra.Command {
//	return &cobra.Command{
//		Use:   "test",
//		Short: "Test subcommand",
//		Long:  "Testing subcommand",
//		Run: func(cmd *cobra.Command, args []string) {
//			fmt.Println("Testing subcommand...")
//		},
//	}
//}

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
	Long:  "print usage in a pretty markdown render using terminal",
	Run: func(cmd *cobra.Command, args []string) {
		markdownrender.RenderUsage()
	},
}
