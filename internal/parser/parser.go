package parser

import (
	"encoding/json"
	"fmt"

	tfjson "github.com/hashicorp/terraform-json"
)

const (
	CREATE string = "create"
	UPDATE string = "update"
	DELETE string = "delete"
)

var (
	resourcesList = make(map[string][]string)
)

func Parser(output []byte) {
	data := tfjson.Plan{}
	if err := json.Unmarshal(output, &data); err != nil {
		panic(err)
	}

	for _, resource := range data.ResourceChanges {
		for _, changes := range resource.Change.Actions {
			action := string(changes)
			resourcesList[action] = append(resourcesList[action], resource.Address)
		}
	}

	PrintResources(fmt.Sprintf("Resource to be created ✅"), resourcesList[CREATE])
	PrintResources(fmt.Sprintf("\nResources to be updated ⚠️"), resourcesList[UPDATE])
	PrintResources(fmt.Sprintf("\nResources to be destroyed ❌"), resourcesList[DELETE])

}

func PrintResources(message string, resources []string) {
	fmt.Println(message)
	for _, resource := range resources {
		fmt.Println(resource)
	}
}
