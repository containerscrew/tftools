package markdownrender

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/charmbracelet/glamour"
)

func ReadUsageFile(version string) string {
	usageURL := fmt.Sprintf("https://raw.githubusercontent.com/containerscrew/tftools/%s/docs/usage.md", version)

	resp, err := http.Get(usageURL)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return string(data)
}

func RenderUsage(version string) {
	data := ReadUsageFile(version)

	out, err := glamour.Render(data, "dark")
	if err != nil {
		log.Println(err)
	}
	fmt.Print(out)
}
