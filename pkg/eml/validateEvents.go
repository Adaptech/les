package eml

import (
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
	for _, event := range stream.Events {
		c.validateEventName(event, boundedContextName, stream)
		c.validateThatEventBelongsToAnAggregate(event, boundedContextName, stream)
		c.validateEventIsTheResultOfACommand(event, boundedContextName, stream)

		// Validate the event's properties:
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
