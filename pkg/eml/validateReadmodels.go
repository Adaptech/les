package eml

import (
	"strings"
)

func (c *Solution) validateReadmodelsIn(context BoundedContext, events map[string]Event, properties map[string]bool) {
	for _, readmodel := range context.Readmodels {
		if len(strings.Trim(readmodel.Readmodel.Name, " ")) == 0 {
			validationError := ValidationError{
				ErrorID: "MissingReadmodelName",
				Context: context.Name,
				Message: "Missing Readmodel Name",
			}
			c.Errors = append(c.Errors, validationError)
		}
		if readmodel.Readmodel.Key == "" {
			validationError := ValidationError{
				ErrorID:   "MissingReadmodelKey",
				Context:   context.Name,
				Readmodel: readmodel.Readmodel.Name,
				Message:   "Missing Readmodel Key",
			}
			c.Errors = append(c.Errors, validationError)
		} else {
			_, propertyExists := properties[readmodel.Readmodel.Key]
			if !propertyExists {
				propertyTypeValidationError := ValidationError{
					ErrorID:   "ReadModelKeyDoesNotExist",
					Context:   context.Name,
					Readmodel: readmodel.Readmodel.Name,
					Message:   "Unknown read model key: '" + readmodel.Readmodel.Key + "' is not a property of any event in '" + context.Name + "'.",
				}
				c.Errors = append(c.Errors, propertyTypeValidationError)

			}
		}
		if len(readmodel.Readmodel.SubscribesTo) == 0 {
			validationError := ValidationError{
				ErrorID:   "MissingReadmodelSubscribesToEvent",
				Context:   context.Name,
				Readmodel: readmodel.Readmodel.Name,
				Message:   "Missing Readmodel SubscribesTo Event(s)",
			}
			c.Errors = append(c.Errors, validationError)
		}
		for _, subscribesToEvent := range readmodel.Readmodel.SubscribesTo {
			_, eventExists := events[subscribesToEvent]
			if !eventExists {
				propertyTypeValidationError := ValidationError{
					ErrorID:   "SubscribesToEventNotFound",
					Context:   context.Name,
					Readmodel: readmodel.Readmodel.Name,
					Message:   "Unknown event: The read model cannot subscribe to an event of type '" + subscribesToEvent + "' because the event type doesn't exist.",
				}
				c.Errors = append(c.Errors, propertyTypeValidationError)
			}
		}
	}

}
