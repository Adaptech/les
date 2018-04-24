package eml

// Ensure that the event is the result of executing a command within the current bounded context
func (c *Solution) validateEventIsTheResultOfACommand(event Event, boundedContextName string, stream Stream) {
	isPostconditionOfCommand := false
	for _, command := range stream.Commands {
		for _, postcondition := range command.Command.Postconditions {
			if postcondition == event.Event.Name {
				isPostconditionOfCommand = true
				break
			}
		}
		if isPostconditionOfCommand {
			break
		}
	}
	if !isPostconditionOfCommand {
		validationError := ValidationError{
			ErrorID: "EventMustBeCommandPostcondition",
			Context: boundedContextName,
			Stream:  stream.Name,
			Message: boundedContextName + " has no command which would result in '" + event.Event.Name + "'.",
		}
		c.Errors = append(c.Errors, validationError)
	}
}
