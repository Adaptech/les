package emd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const dotTemplate string = `
digraph G {
	graph [ bgcolor=transparent style=filled ] { 
		subgraph commands {
            node [shape=box pos=1 fontname="Helvetica bold" style=filled fillcolor="#75a6f1"]
            { rank=same;
%s
            }
        }
        subgraph events {
            node [shape=box fontname="Helvetica bold" style=filled fillcolor="#f0dc97" ]
            { rank=same;
%s
            }
        }

        subgraph documents {
            node [shape=box fontname="Helvetica bold" style=filled fillcolor="#67bfa2" ]
            { rank=same; 
%s
            }
        }
%s
	}	
}
`
const commandTemplate = "                %s[label=\"{ %s%s}\" shape = \"record\"]\n"
const commandToCommandArrowTemplate = "                %s->%s[style=invis]\n"
const commandEventArrowTemplate = "                %s->%s\n"

const eventTemplate = "                %s[label=\"{ %s%s}\" shape = \"record\"]\n"
const eventToEventArrowTemplate = "                %s->%s[style=invis]\n"
const eventDocumentArrowTemplate = "                %s->%s [style=invis]\n"

const documentTemplate = "                %s[label=\"{ %s%s}\" shape = \"record\"]\n"
const documentToDocumentArrowTemplate = "                %s->%s[style=invis]"

const parameterTemplate = "&#92; &#92; &#92; -&#92; %s\\l"

const invisibleSpacerTemplate = "                %s%v[style=invis]\n                %s->%s%v [style=invis]\n"

// ToGraphViz turns EMD into a GraphViz digraph which can be used to generate images of the event storming described by the EMD input
func ToGraphViz(emdString string) string {
	postitNotes := strings.Split(emdString, "\n")
	commandList := ""
	eventList := ""
	documentList := ""
	connectionsList := ""
	var previousCommand string
	var previousEvent string
	var previousDocument string
	writeCommandEventArrow := false
	nonWordCharacterRemover := regexp.MustCompile("\\W")
	for index, postitNote := range postitNotes {
		// This is a hack to support commands with parameters without the '//' slashes.
		// "command -> // param1" and "command -> param1" should both be valid ... because
		// it turns out that users commonly forget the forward slashes when entering EMD commands.
		if strings.Contains(postitNote, "->") && !strings.Contains(postitNote, "//") {
			postitNote = strings.Replace(postitNote, "->", "-> // ", -1)
		}
		graphVizDotPropertiesList := ""
		// Guard against DOS attacks by limiting length of lines and number of lines:
		if len(postitNote) > 200 {
			continue
		}
		if index > 300 || len(postitNote) > 200 {
			break
		}
		// Ignore lines full of spaces
		if len(strings.Replace(postitNote, " ", "", -1)) == 0 {
			continue
		}
		// Ignore lines which aren't commands but have the '>'
		if strings.Contains(postitNote, ">") && !strings.Contains(postitNote, "->") {
			continue
		}
		// Ignore single character & blank lines
		if len(postitNote) <= 2 {
			continue
		}
		// Skip comment lines:
		if string(postitNote[0]) == "#" {
			continue
		}
		postitNoteParts := strings.Split(postitNote, "//")
		name := postitNoteParts[0]
		name = strings.Trim(name, " ")

		// Determine if there are any parameters, e.g.
		// UserRegistered // userId,firstName,lastName
		if len(postitNoteParts) > 1 {
			properties := postitNoteParts[1]
			if len(properties) > 0 {
				propertiesList := strings.Split(properties, ",")
				for _, property := range propertiesList {
					property = strings.Trim(property, " ")
					const maximumPropertyLength = 100
					if len(property) > maximumPropertyLength {
						property = property[0:maximumPropertyLength]
					}
					property = nonWordCharacterRemover.ReplaceAllString(property, "")
					if len(property) > 0 {
						graphVizDotProperty := fmt.Sprintf(parameterTemplate, property)
						if len(graphVizDotPropertiesList) == 0 {
							graphVizDotPropertiesList = " | "
						}
						graphVizDotPropertiesList = graphVizDotPropertiesList + graphVizDotProperty
					}
				}
			}
		}

		lastTwoChars := name[len(name)-2 : len(name)]
		lastChar := name[len(name)-1 : len(name)]
		if lastTwoChars == "->" {
			// Generate DOT for command, escaping double quotes:
			commandLabel := strings.Replace(name[0:len(name)-2], "\"", "\\\"", -1)
			command := nonWordCharacterRemover.ReplaceAllString(commandLabel, "") + strconv.Itoa(index)

			var dotCommand string
			dotCommand = fmt.Sprintf(commandTemplate, command, commandLabel, graphVizDotPropertiesList)

			var dotCommandToCommandArrow string
			if len(previousCommand) > 0 {
				dotCommandToCommandArrow = fmt.Sprintf(commandToCommandArrowTemplate, previousCommand, command)
			} else {
				dotCommandToCommandArrow = ""
			}
			previousCommand = command
			commandList = commandList + dotCommand + dotCommandToCommandArrow
			writeCommandEventArrow = true
		} else if lastChar == "*" {
			// Generate DOT for document, escaping double quotes:
			documentLabel := strings.Replace(name[0:len(name)-1], "\"", "\\\"", -1)

			// Remove all non-word characters:
			document := nonWordCharacterRemover.ReplaceAllString(documentLabel, "") + strconv.Itoa(index)
			var dotDocument string
			dotDocument = fmt.Sprintf(documentTemplate, document, documentLabel, graphVizDotPropertiesList)

			var dotDocumentToDocumentArrow string
			if len(previousDocument) > 0 {
				dotDocumentToDocumentArrow = fmt.Sprintf(documentToDocumentArrowTemplate, previousDocument, document)
			} else {
				dotDocumentToDocumentArrow = ""
			}
			previousDocument = document
			documentList = documentList + dotDocument + dotDocumentToDocumentArrow

			connectionArrow := fmt.Sprintf(eventDocumentArrowTemplate, previousEvent, document)
			connectionsList = connectionsList + connectionArrow
		} else {
			// Generate DOT for event, escaping double quotes:
			eventLabel := strings.Replace(name, "\"", "\\\"", -1)

			event := nonWordCharacterRemover.ReplaceAllString(eventLabel, "") + strconv.Itoa(index)

			var dotEvent string
			dotEvent = fmt.Sprintf(eventTemplate, event, eventLabel, graphVizDotPropertiesList)
			var invisibleSpacer string
			var dotEventToEventArrow string
			if len(previousEvent) > 0 {
				dotEventToEventArrow = fmt.Sprintf(eventToEventArrowTemplate, previousEvent, event)
			} else {
				dotEventToEventArrow = ""
			}
			previousEvent = event
			eventList = eventList + dotEvent + dotEventToEventArrow

			if len(previousCommand) > 0 {
				if writeCommandEventArrow {
					connectionArrow := fmt.Sprintf(commandEventArrowTemplate, previousCommand, event)
					connectionsList = connectionsList + connectionArrow
					writeCommandEventArrow = false
				} else {
					invisibleSpacer = fmt.Sprintf(invisibleSpacerTemplate, event, index, previousCommand, event, index)
					commandList = commandList + invisibleSpacer
					previousCommand = fmt.Sprintf("%v%v", event, index)
				}
			}
		}
	}
	digraph := fmt.Sprintf(dotTemplate, commandList, eventList, documentList, connectionsList)
	return digraph
}
