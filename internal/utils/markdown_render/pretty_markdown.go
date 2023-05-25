package markdownrender

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/charmbracelet/glamour"
)

const (
	usageURL = "https://raw.githubusercontent.com/containerscrew/tftools/main/docs/usage.md"
)

func ReadUsageFile() string {
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

func RenderUsage() {
	data := ReadUsageFile()

	out, err := glamour.Render(data, "dark")
	if err != nil {
		log.Println(err)
	}
	fmt.Print(out)
}
