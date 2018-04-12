package eml

import (
	"fmt"
	"regexp"
	"strings"
)

// Validate if the system is an event sourced system suitable for generating an API.
func (c *Solution) Validate() {
	errors := []ValidationError{}
	c.Errors = errors
	if len(c.Name) == 0 {
		validationError := ValidationError{ErrorID: "NoSolutionName", Message: "No Solution Name"}
		c.Errors = append(c.Errors, validationError)
	}
	if len(c.Contexts) == 0 {
		validationError := ValidationError{ErrorID: "NoBoundedContext", Message: "No Bounded Context"}
		c.Errors = append(c.Errors, validationError)
	}

	for _, context := range c.Contexts {
		nameCounts := streamCommandEventAndReadmodelNameCountsIn(context)
		for name, noOfOccurrences := range nameCounts {
			if noOfOccurrences > 1 {
				message := fmt.Sprintf("Names must be unique within a bounded context: '%v' found %v times in '%v'.", name, noOfOccurrences, context.Name)
				validationError := ValidationError{ErrorID: "NamesMustBeUnique", Message: message}
				c.Errors = append(c.Errors, validationError)
			}
		}
		eventLookup := allEventsInBoundedContext(context)
		c.validateStreamsIn(context)
		c.validateReadmodelsIn(context, eventLookup)
	}
}

func allEventsInBoundedContext(context BoundedContext) map[string]Event {
	eventMap := make(map[string]Event)
	for _, stream := range context.Streams {
		for _, event := range stream.Events {
			eventMap[event.Event.Name] = event
		}
	}
	return eventMap
}

func streamCommandEventAndReadmodelNameCountsIn(context BoundedContext) map[string]int {
	names := make(map[string]int)
	for _, stream := range context.Streams {
		names[stream.Name] = names[stream.Name] + 1
		for _, command := range stream.Commands {
			names[command.Command.Name] = names[command.Command.Name] + 1
		}
		for _, event := range stream.Events {
			names[event.Event.Name] = names[event.Event.Name] + 1
		}
	}

	for _, readmodel := range context.Readmodels {
		names[readmodel.Readmodel.Name] = names[readmodel.Readmodel.Name] + 1
	}

	return names
}

func isValidName(name string) bool {
	var isAllAlphanumericAndSpacesAndNumbersWithoutNonAlphaAtStartAndEnd = regexp.MustCompile(`(?m)^[A-Za-z][\w]*[\s]*[\w ]*[A-Za-z]$`)
	if strings.Contains(name, " ") {
		return false
	}
	isValid := isAllAlphanumericAndSpacesAndNumbersWithoutNonAlphaAtStartAndEnd.MatchString(name)
	return isValid
}

func isValidCommandName(name string) bool {
	var startsWithAlpha = regexp.MustCompile(`^[a-zA-Z][\s\w]*`)
	return startsWithAlpha.MatchString(name)
}
