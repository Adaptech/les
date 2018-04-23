package convert_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/convert"
	"github.com/Adaptech/les/pkg/emd"
)

func TestShouldCreateEvents(t *testing.T) {
	input := []string{
		"# Timesheets",
		"User Registered",
		"Timesheet Created",
		"Timesheet Submitted",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	streams := markup.Eml.Contexts[0].Streams
	hasUserAggregate := false
	hasTimesheetAggregate := false
	for _, stream := range streams {
		if stream.Name == "User" && !hasUserAggregate {
			hasUserAggregate = true
		}
		if stream.Name == "Timesheet" && !hasTimesheetAggregate {
			hasTimesheetAggregate = true
		}
		if !(stream.Name == "Timesheet") && !(stream.Name == "User") {
			t.Error("expected only Timesheet and User streams - " + stream.Name + " unknown.")
		}
	}
	if !hasUserAggregate {
		t.Error("expected a user aggregate stream")
	}
	if !hasTimesheetAggregate {
		t.Error("expected a timesheet aggregate stream")
	}

	if len(streams[0].Events) == 0 {
		t.Error("expected different no of User events.")
	}
	if len(streams[1].Events) == 0 {
		t.Error("expected different no of Timesheet events.")
	}
}
func TestEventsShouldHaveCorrectNumberAndTypeOfProperties(t *testing.T) {
	input := []string{
		"# Timesheets",
		"User Registered // userId, password",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	streams := markup.Eml.Contexts[0].Streams
	if len(streams[0].Events) != 1 {
		t.Error("expected different no of User events.")
	} else {
		event := streams[0].Events[0]
		if len(event.Event.Properties) != 2 {
			t.Error("expected different no of User Registered event properties.")
		} else {
			userIdProperty := event.Event.Properties[0]
			passwordProperty := event.Event.Properties[1]
			if userIdProperty.Name != "userId" {
				t.Error("expected different event property name")
			}
			if userIdProperty.Type != "string" {
				t.Error("expected different event property type")
			}
			if userIdProperty.IsHashed != false {
				t.Error("expected different event property encryption setting")
			}

			if passwordProperty.Name != "password" {
				t.Error("expected different event property name")
			}
			if passwordProperty.Type != "string" {
				t.Error("expected different event property type")
			}
			if passwordProperty.IsHashed != true {
				t.Error("expected different event property encryption setting")
			}
		}
	}
}

func TestEventsWithoutAggregateIdShouldHaveAggregateIdAdded(t *testing.T) {
	input := []string{
		"# Timesheets",
		"User Registered // password",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	streams := markup.Eml.Contexts[0].Streams
	if len(streams[0].Events) != 1 {
		t.Error("expected different no of User events.")
	} else {
		event := streams[0].Events[0]
		if len(event.Event.Properties) != 2 {
			t.Error("expected different no of User Registered event properties.")
		} else {
			userIdProperty := event.Event.Properties[1]
			if userIdProperty.Name != "userId" {
				t.Error("expected different event property name")
			}
			if userIdProperty.Type != "string" {
				t.Error("expected different event property type")
			}
			if userIdProperty.IsHashed != false {
				t.Error("expected different event property encryption setting")
			}
		}
	}
}

func TestShouldNotAllowOnOneWordEvents(t *testing.T) {
	input := []string{
		"# Timesheets",
		"UserRegistered",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	if len(markup.MarkdownValidationErrors) != 1 {
		t.Error("expected validation error")
	}
}

func TestShouldStripSpacesFromEventPropertyNames(t *testing.T) {
	input := []string{
		"# Timesheets",
		"User Registered // User Email",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	if len(markup.MarkdownValidationErrors) != 0 {
		t.Error("expected no validation error")
	}

	propertyName := markup.Eml.Contexts[0].Streams[0].Events[0].Event.Properties[0].Name

	if propertyName != "UserEmail" {
		t.Error("expected property name to have no spaces.")
	}
}

func TestShouldStripSpacesFromCommandParameterNames(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register User->//User Email",
		"User Registered // email",
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	if len(markup.MarkdownValidationErrors) != 0 {
		t.Error("expected no validation error")
	}

	parameterName := markup.Eml.Contexts[0].Streams[0].Commands[0].Command.Parameters[0].Name

	if parameterName != "UserEmail" {
		t.Error("expected command parameter name to have no spaces.")
	}
}
