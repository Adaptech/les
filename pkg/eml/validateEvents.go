package eml

import (
	"regexp"
	"strings"
)

func (c *Solution) validateEvents(boundedContextName string, stream Stream) {
	if len(stream.Events) == 0 {
		validationError := ValidationError{
			ErrorID: "NoEvents",
			Context: boundedContextName,
			Stream:  stream.Name,
			Message: "Event stream '" + stream.Name + "': Must have at least one event that is being published to it.",
		}
		c.Errors = append(c.Errors, validationError)
	}
	var isAllAlphanumericAndSpacesAndNumbersWithoutNonAlphaAtStartAndEnd = regexp.MustCompile(`(?m)^[A-Za-z][\w]*[\s]*[\w ]*[A-Za-z]$`)
	for _, event := range stream.Events {
		// Verify that the event name follows rules which make it suitable for variable naming when generating code in most programming languages.
		// Spaces are OK for readability - they are intended to be stripped out in code generators when the name is used as a variable name.
		if !isAllAlphanumericAndSpacesAndNumbersWithoutNonAlphaAtStartAndEnd.MatchString(event.Event.Name) {
			validationError := ValidationError{
				ErrorID: "InvalidEventName",
				Context: boundedContextName,
				Stream:  stream.Name,
				Message: event.Event.Name + ": Invalid event name.",
			}
			c.Errors = append(c.Errors, validationError)
		}

		if len(strings.Trim(event.Event.Name, " ")) == 0 {
			validationError := ValidationError{
				ErrorID: "NoEventID",
				Context: boundedContextName,
				Stream:  stream.Name,
				Message: event.Event.Name + ": Event IDs cannot be blank.",
			}
			c.Errors = append(c.Errors, validationError)
		}

		// Verify that the event has an aggregate/stream ID so it's known what aggregate/stream's state change it resulted from:
		hasAggregateID := false
		for _, property := range event.Event.Properties {
			if strings.ToLower(property.Name) == strings.ToLower(stream.Name+"Id") {
				hasAggregateID = true
				break
			}
		}
		if !hasAggregateID {
			validationError := ValidationError{
				ErrorID: "NoAggregateId",
				Context: boundedContextName,
				Stream:  stream.Name,
				Message: "Missing stream/aggregate ID property for  '" + event.Event.Name + "' event.",
			}
			c.Errors = append(c.Errors, validationError)
		}

		for _, property := range event.Event.Properties {
			// Verify that properties have names:
			if len(strings.Trim(property.Name, " ")) == 0 {
				propertyNameValidationError := ValidationError{
					ErrorID: "NoPropertyName",
					Context: boundedContextName,
					Stream:  stream.Name,
					Message: event.Event.Name + ": Event property names cannot be blank.",
				}
				c.Errors = append(c.Errors, propertyNameValidationError)
			}

			// Verify that properties have valid names
			if !isValidName(property.Name) {
				parameterNameError := ValidationError{
					ErrorID: "InvalidEventPropertyName",
					Context: boundedContextName,
					Stream:  stream.Name,
					Message: "Invalid event property name '" + property.Name + "': Names must start with a character and contain no spaces.",
				}
				c.Errors = append(c.Errors, parameterNameError)
			}

			// Verify that properties have property types:
			if len(strings.Trim(property.Type, " ")) == 0 {
				propertyTypeValidationError := ValidationError{
					ErrorID: "NoPropertyType",
					Context: boundedContextName,
					Stream:  stream.Name,
					Message: event.Event.Name + "Event property types cannot be blank: ",
				}
				c.Errors = append(c.Errors, propertyTypeValidationError)
			}
		}

	}
}
