package convert_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/convert"
	"github.com/Adaptech/les/pkg/emd"
	"github.com/Adaptech/les/pkg/eml"
)

func TestShouldCreateReadmodel(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Create ->",
		"User Created // name",
		"User List* // name",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	conversionResult, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	context := conversionResult.Eml.Contexts[0]
	if len(conversionResult.Eml.Errors) > 0 {
		t.Error("unexpected EML validation error")
	}
	if len(conversionResult.MarkdownValidationErrors) > 0 {
		t.Error("unexpected EMD validation error")
	}

	if len(context.Readmodels) != 1 {
		t.Error("expected different no of readmodels.")
	} else {
		readmodel := context.Readmodels[0]
		if readmodel.Readmodel.Name != "UserList" {
			t.Error("Expected different readmodel ID")
		}
	}
}

func TestShouldFailValidationIfNoReadmodelKeyPresent(t *testing.T) {
	input := []string{
		"# Timesheets",
		"User Registered // userId, email",
		"User List*",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	context := markup.Eml.Contexts[0]

	if len(context.Readmodels) != 1 {
		t.Error("expected different no of readmodels.")
	}
	if !hasError("MissingReadmodelKey", markup.Eml.Errors) {
		t.Error("expected error")
	}
}

func TestReadmodelShouldHaveSubscribesToEvents(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register User->",
		"User Registered // userId, email, password",
		"Create Timesheet->",
		"Timesheet Created // timesheetId, userId, date",
		"User List* // userId, email, date",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	context := markup.Eml.Contexts[0]

	if len(markup.Eml.Errors) != 0 || len(markup.MarkdownValidationErrors) != 0 {
		t.Error("expected no eml or emd validation errors")
	}

	readmodel := context.Readmodels[0]
	if len(readmodel.Readmodel.SubscribesTo) != 2 {
		t.Error("expected different number of readmodel subscribesTo events.")
	} else {

		hasUserRegisteredEvent := false
		hasTimesheetCreatedEvent := false
		for _, event := range readmodel.Readmodel.SubscribesTo {
			if event == "UserRegistered" && !hasUserRegisteredEvent {
				hasUserRegisteredEvent = true
			}
			if event == "TimesheetCreated" && !hasTimesheetCreatedEvent {
				hasTimesheetCreatedEvent = true
			}
			if !(event == "TimesheetCreated") && !(event == "UserRegistered") {
				t.Error("expected only Timesheet and User events - " + event + " unknown.")
			}
		}
		if !hasUserRegisteredEvent {
			t.Error("expected a user aggregate stream")
		}
		if !hasTimesheetCreatedEvent {
			t.Error("expected a timesheet aggregate stream")
		}
	}
}

func TestReadmodelShouldNotSubscribeToUnneededEvents(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register User->",
		"User Registered // userId, email, password",
		"Create Timesheet->",
		"Timesheet Created // timesheetId, date",
		"Email List* // userId, email",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	context := markup.Eml.Contexts[0]

	if len(markup.Eml.Errors) != 0 || len(markup.MarkdownValidationErrors) != 0 {
		t.Error("expected no eml or emd validation errors")
	}

	readmodel := context.Readmodels[0]
	if len(readmodel.Readmodel.SubscribesTo) != 1 {
		t.Error("expected different number of readmodel subscribesTo events.")
	} else {
		firstEvent := readmodel.Readmodel.SubscribesTo[0]
		if firstEvent != "UserRegistered" {
			t.Error("expected different readmodel subscribesTo event")
		}
	}
}

// Subscribing to all events which have stream ids (userId, timesheetId, ...) will lead to
// read models which try to subscribe to all events which have the stream/aggregate ID and will contain all
// their properties in a denormalized flat table. This is not a desirable behaviour at this time.
func TestReadmodelShouldNotSubscribeToEventsDueToStreamIdFields(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register User->",
		"User Registered // userId, email, password",
		"Create Timesheet->",
		"Timesheet Created // timesheetId, date",
		"Email List* // userId, email, timesheetId",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	context := markup.Eml.Contexts[0]

	if len(markup.Eml.Errors) != 0 || len(markup.MarkdownValidationErrors) != 0 {
		t.Error("expected no eml or emd validation errors")
	}

	readmodel := context.Readmodels[0]
	if len(readmodel.Readmodel.SubscribesTo) != 1 {
		t.Error("expected different number of readmodel subscribesTo events.")
	} else {
		firstEvent := readmodel.Readmodel.SubscribesTo[0]
		if firstEvent != "UserRegistered" {
			t.Error("expected different readmodel subscribesTo event")
		}
	}
}

// For a EMD read model to be valid it needs to have at least one property because the first property
// after the slashes is assumed to be the read model key - which is required because we can't have read models without keys.
// Example:  'UserList* // userId' is valid. 'UserList* //' or 'UserList*' is not.
func TestReadmodelShouldHaveMissingKeyValidationErrorIfThereIsNoPropertyAfterDoubleSlashes(t *testing.T) {
	input := []string{
		"# Users",
		"Register User->",
		"User Registered // userId, email, password",
		"UserList* //",
	}

	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	context := markup.Eml.Contexts[0]

	if len(markup.Eml.Errors) != 2 || len(markup.MarkdownValidationErrors) != 0 {
		t.Error("expected different numbers of eml and emd validation errors")
	}

	readmodel := context.Readmodels[0]
	if readmodel.Readmodel.Key != "" {
		t.Error("Expected a read model without a key (because no parameters were specified).")
	}
	if !hasError("MissingReadmodelKey", markup.Eml.Errors) {
		t.Error("expected error")
	}
	if !hasError("MissingReadmodelSubscribesToEvent", markup.Eml.Errors) {
		t.Error("expected error")
	}
}

func hasError(errorId string, errors []eml.ValidationError) bool {
	for _, err := range errors {
		if err.ErrorID == errorId {
			return true
		}
	}
	return false
}
