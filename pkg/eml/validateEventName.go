package eml

import (
	"regexp"
	"strings"
)

var isAllAlphanumericAndSpacesAndNumbersWithoutNonAlphaAtStartAndEnd = regexp.MustCompile(`(?m)^[A-Za-z][\w]*[\s]*[\w ]*[A-Za-z]$`)

// Verify that the event name follows rules which make it suitable for variable naming when generating code in most programming languages.
// Spaces are OK for readability - they are intended to be stripped out in code generators when the name is used as a variable name.
func (c *Solution) validateEventName(event Event, boundedContextName string, stream Stream) {
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
}
