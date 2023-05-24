package markdownrender

import (
	"embed"
	"fmt"
	"log"

	"github.com/charmbracelet/glamour"
)

// Embed usage documentation.
//go:generate cp -r ../../../docs/usage.md ./usage.md
//go:embed usage.md

var usage embed.FS

func RenderUsage() {
	usage, err := usage.ReadFile("docs/usage.md")
	if err != nil {
		log.Fatal(err)
	}

	out, err := glamour.Render(string(usage), "dark")
	if err != nil {
		log.Println(err)
	}
	fmt.Print(out)
}
