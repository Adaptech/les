package eml

import (
	"strings"
)

// Verify that the event has an aggregate/stream ID so it's known what aggregate/stream's state change it resulted from:
func (c *Solution) validateThatEventBelongsToAnAggregate(event Event, boundedContextName string, stream Stream) {
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
}
