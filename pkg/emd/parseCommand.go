package emd

import (
	"regexp"
	"strings"
)

func parseCommand(emdInput string, lineItems []Item) []Item {
	if strings.Contains(emdInput, "//") {
		re, err := regexp.Compile("^(.*) *-> *\\/\\/ *(.*)")
		if err != nil {
			panic(err)
		}
		command := re.FindAllStringSubmatch(emdInput, -1)
		if len(command) > 0 {
			var parameters []Parameter
			first := command[0]
			propertiesList := first[2]
			propertiesList = strings.Trim(propertiesList, ", ")
			inputProperties := strings.Split(propertiesList, ",")
			for _, inputParameter := range inputProperties {
				var parsedParameter = Parameter{Name: strings.Trim(inputParameter, " ")}
				parameters = append(parameters, parsedParameter)
			}

			lineItems = append(lineItems, Command{Name: strings.Trim(command[0][1], " "), Parameters: parameters})
		}
	} else {
		spacesRemoved := strings.Trim(emdInput, " ")
		commandName := strings.Trim(strings.Replace(spacesRemoved, "->", "", -1), " ")
		lineItems = append(lineItems, Command{Name: commandName})
	}
	return lineItems
}
