package emd

import (
	"regexp"
	"strings"
)

func parseDocument(emdInput string, lineItems []Item) []Item {
	re, err := regexp.Compile("^(.*) *\\* (.*)")
	if err != nil {
		panic(err)
	}
	emdInput = strings.Replace(emdInput, "//", "", -1)
	document := re.FindAllStringSubmatch(emdInput, -1)
	if len(document) > 0 {
		var properties []Property
		first := document[0]
		propertiesList := first[2]
		propertiesList = strings.Trim(propertiesList, ", ")
		inputProperties := strings.Split(propertiesList, ",")
		for _, inputParameter := range inputProperties {
			if inputParameter != "" {
				var parsedParameter = Property{Name: strings.Trim(inputParameter, " ")}
				properties = append(properties, parsedParameter)
			}
		}
		lineItems = append(lineItems, Document{Name: strings.Trim(document[0][1], " "), Properties: properties})
		return lineItems
	}

	return lineItems
}
