package emd

import (
	"regexp"
	"strings"
)

func parseEvent(emdInput string, lineItems []Item) []Item {
	if strings.Contains(emdInput, "//") {
		re, err := regexp.Compile("^(.*) *\\/\\/ *(.*)")
		if err != nil {
			panic(err)
		}
		event := re.FindAllStringSubmatch(emdInput, -1)
		if len(event) > 0 {
			var properties []Property
			first := event[0]
			propertiesList := first[2]
			propertiesList = strings.Trim(propertiesList, ", ")
			inputProperties := strings.Split(propertiesList, ",")
			for _, inputProperty := range inputProperties {
				var parsedProperty = Property{Name: strings.Trim(inputProperty, " ")}
				properties = append(properties, parsedProperty)
			}

			lineItems = append(lineItems, Event{Name: strings.Trim(event[0][1], " "), Properties: properties})
		}
	} else {
		lineItems = append(lineItems, Event{Name: strings.Trim(emdInput, " ")})
	}
	return lineItems
}
