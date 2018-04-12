package emd

import (
	"strings"
)

// Parse Event Markdown to Event Markup Language
func Parse(emdInput []string) (Emd, error) {
	eventMarkdown := Emd{}
	var lineItems = []Item{}
	for _, line := range emdInput {
		if len(line) <= 2 || line[:2] == "//" {
			continue
		}
		if strings.Contains(line, "#") {
			lineItems = parseComment(line, lineItems)
		} else if !strings.Contains(line, "->") && !strings.Contains(line, "*") {
			lineItems = parseEvent(line, lineItems)
		} else if strings.Contains(line, "->") {
			lineItems = parseCommand(line, lineItems)
		} else if strings.Contains(line, "*") {
			lineItems = parseDocument(line, lineItems)
		}
	}
	eventMarkdown.Lines = lineItems
	return eventMarkdown, nil
}
