package parser

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	tfjson "github.com/hashicorp/terraform-json"
)

const (
	CREATE string = "create"
	UPDATE string = "update"
	DELETE string = "delete"
	TAG    string = "tag"
	NOOP   string = "no-op"
)

var (
	resourcesList = make(map[string][]string)
)

func Parser(output []byte, showTags, showUnchanged, compact, useMarkdown bool) {
	var data tfjson.Plan
	if err := json.Unmarshal(output, &data); err != nil {
		fmt.Printf("Error unmarshalling plan: %v\n", err)
		return
	}

	for _, resourceChange := range data.ResourceChanges {
		processResourceChange(resourceChange, showTags)
	}

	PrintPlanSummary(showTags, showUnchanged, compact, useMarkdown)
}

func processResourceChange(resourceChange *tfjson.ResourceChange, showTags bool) {
	isUpdate := contains(resourceChange.Change.Actions, tfjson.ActionUpdate)

	if isUpdate {
		// Check if it's only a tag change
		isOnlyTagChange, err := checkOnlyTagChanges(resourceChange)
		if err != nil {
			fmt.Printf("Error checking for only tag changes: %v\n", err)
			return
		}
		if isOnlyTagChange && showTags {
			// Only add to TAG if it's only a tag change
			resourcesList[TAG] = append(resourcesList[TAG], resourceChange.Address)
			return
		}

		isTagChange := hasTagChanges(resourceChange)
		if err != nil {
			fmt.Printf("Error checking for tag changes: %v\n", err)
			return
		}

		if isTagChange && showTags {
			resourcesList[TAG] = append(resourcesList[TAG], resourceChange.Address)
			resourcesList[UPDATE] = append(resourcesList[UPDATE], resourceChange.Address)
			return
		}

		addActionToResourceList(resourceChange.Change.Actions, resourceChange.Address)

	} else {
		// Not an update, add to other categories as necessary
		addActionToResourceList(resourceChange.Change.Actions, resourceChange.Address)
	}
}

func hasTagChanges(resourceChange *tfjson.ResourceChange) bool {
	beforeRaw, err := json.Marshal(resourceChange.Change.Before)
	if err != nil {
		fmt.Printf("Error marshaling before state: %v\n", err)
		return false
	}
	afterRaw, err := json.Marshal(resourceChange.Change.After)
	if err != nil {
		fmt.Printf("Error marshaling after state: %v\n", err)
		return false
	}

	var beforeMap, afterMap map[string]interface{}
	if err := json.Unmarshal(beforeRaw, &beforeMap); err != nil {
		fmt.Printf("Error unmarshaling before state: %v\n", err)
		return false
	}
	if err := json.Unmarshal(afterRaw, &afterMap); err != nil {
		fmt.Printf("Error unmarshaling after state: %v\n", err)
		return false
	}

	if equal(beforeMap["tags"], afterMap["tags"]) && equal(beforeMap["tags_all"], afterMap["tags_all"]) {
		return false
	}

	return true
}

func addActionToResourceList(actions []tfjson.Action, address string) {
	for _, action := range actions {
		resourcesList[string(action)] = append(resourcesList[string(action)], address)
	}
}

func PrintResources(message string, resources []string, bulletSymbol string, color *color.Color, compact bool, useMarkdown bool) {
	if len(resources) != 0 {
		if useMarkdown {
			fmt.Printf("## %s\n\n", message) // Markdown header for the message
			for _, resource := range resources {
				var emoji string
				switch bulletSymbol {
				case "+":
					emoji = "✅" // Green check mark for create
				case "~":
					emoji = "⚠️" // Yellow warning sign for update
				case "-":
					emoji = "🧨" // Red circle for destroy
				case "#":
					emoji = "#️⃣" // Blue diamond for tag/untag
				case "•":
					emoji = "🔷" // Blue circle for unchanged
				default:
					emoji = "➡️" // Default arrow
				}
				fmt.Printf("* %s %s\n", emoji, resource)
			}
		} else {
			fmt.Println(message)
			for _, resource := range resources {
				color.Printf("  %s ", bulletSymbol)
				fmt.Println(resource)
			}
		}
		if !compact {
			fmt.Println()
		}
	}
}

func PrintPlanSummary(showTags, showUnchanged, compact, useMarkdown bool) {
	if showUnchanged {
		PrintResources("🔵 Unchanged:", resourcesList[NOOP], "•", color.New(color.FgBlue), compact, useMarkdown)
	}
	if showTags {
		PrintResources("🟣 Tag/Untag:", resourcesList[TAG], "#", color.New(color.FgMagenta), compact, useMarkdown)
	}
	PrintResources("🟢 Create:", resourcesList[CREATE], "+", color.New(color.FgGreen), compact, useMarkdown)
	PrintResources("🟡 Update:", resourcesList[UPDATE], "~", color.New(color.FgYellow), compact, useMarkdown)
	PrintResources("🔴 Destroy:", resourcesList[DELETE], "-", color.New(color.FgRed), compact, useMarkdown)
}

func checkOnlyTagChanges(resourceChange *tfjson.ResourceChange) (bool, error) {
	beforeRaw, err := json.Marshal(resourceChange.Change.Before)
	if err != nil {
		return false, fmt.Errorf("failed to marshal before state: %v", err)
	}
	afterRaw, err := json.Marshal(resourceChange.Change.After)
	if err != nil {
		return false, fmt.Errorf("failed to marshal after state: %v", err)
	}

	var beforeMap, afterMap map[string]interface{}
	if err := json.Unmarshal(beforeRaw, &beforeMap); err != nil {
		return false, fmt.Errorf("failed to unmarshal before state: %v", err)
	}
	if err := json.Unmarshal(afterRaw, &afterMap); err != nil {
		return false, fmt.Errorf("failed to unmarshal after state: %v", err)
	}

	if equal(beforeMap, afterMap) {
		return false, nil
	}

	for key := range beforeMap {
		if key != "tags" && key != "tags_all" {
			if vAfter, exists := afterMap[key]; exists {
				if !equal(beforeMap[key], vAfter) {
					return false, nil
				}
			} else {
				return false, nil
			}
		}
	}

	for key := range afterMap {
		if key != "tags" && key != "tags_all" {
			if vBefore, exists := beforeMap[key]; exists {
				if !equal(vBefore, afterMap[key]) {
					return false, nil
				}
			} else {
				return false, nil
			}
		}
	}

	return true, nil
}

func equal(a, b interface{}) bool {
	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)
	return string(aJSON) == string(bJSON)
}

func contains(slice []tfjson.Action, val tfjson.Action) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
