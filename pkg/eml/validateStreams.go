package eml

func (c *Solution) validateStreamsIn(context BoundedContext) {
	if len(context.Streams) == 0 {
		validationError := ValidationError{
			ErrorID: "NoStreams",
			Context: context.Name,
			Message: "No event streams found",
		}
		c.Errors = append(c.Errors, validationError)
		return
	}
	for _, stream := range context.Streams {
		eventLookup := allEventsIn(stream)
		readmodels := allReadmodelsInBoundedContext(context)
		c.validateCommands(context.Name, stream, eventLookup, readmodels)
		c.validateEvents(context.Name, stream)
	}
}

func allEventsIn(stream Stream) map[string]Event {
	eventMap := make(map[string]Event)
	for _, event := range stream.Events {
		eventMap[event.Event.Name] = event
	}
	return eventMap
}

func allReadmodelsInBoundedContext(context BoundedContext) map[string]Readmodel {
	readmodelMap := make(map[string]Readmodel)
	for _, readModel := range context.Readmodels {
		readmodelMap[readModel.Readmodel.Name] = readModel
	}
	return readmodelMap
}
