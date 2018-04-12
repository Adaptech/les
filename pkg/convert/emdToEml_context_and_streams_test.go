package convert_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/convert"
	"github.com/Adaptech/les/pkg/emd"
)

func TestShouldCreateBoundedContextFromFirstComment(t *testing.T) {
	input := []string{"# Timesheets"}
	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	markup, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	if len(markup.Eml.Contexts) != 1 {
		t.Error("expected different number of contexts")
		return
	}
	if markup.Eml.Contexts[0].Name != "Timesheets" {
		t.Error("expected different markup.Eml.Contexts.Name")
	}
}

func TestShouldCreateStreamsFromFirstWordsInEventNames(t *testing.T) {
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
}

// This is because by convention the first word ("User") of an event is the name of the stream the event belongs to.
func TestShouldNotCreateStreamsFromOneWordEvents(t *testing.T) {
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

	streams := markup.Eml.Contexts[0].Streams
	if len(streams) != 0 {
		t.Error("expected different number of Contexts.Streams")
	}
	if len(markup.MarkdownValidationErrors) != 1 {
		t.Error("expected an error")
	} else {
		if markup.MarkdownValidationErrors[0].ErrorID != "NoStreamName" {
			t.Error("expected different error")
		}
	}
}
