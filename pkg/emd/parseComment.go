package emd

import (
	"regexp"
	"strings"
)

func parseComment(emdInput string, lineItems []Item) []Item {
	re, _ := regexp.Compile("^ *# *(.*)")
	comment := re.FindAllStringSubmatch(emdInput, -1)
	if len(comment) > 0 {
		lineItems = append(lineItems, Comment{Text: strings.Trim(comment[0][1], " ")})
	}
	return lineItems
}
