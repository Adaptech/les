package emd_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/emd"
)

func TestShouldGetEventWithProperties(t *testing.T) {
	input := []string{"User Registered // userId, email, password                     "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Event:
		if result.Lines[0].(emd.Event).Name != "User Registered" {
			t.Error("Unexpected Event.Name")
		}
	default:
		t.Error("expected event")
	}
}

func TestShouldGetProperties(t *testing.T) {
	input := []string{"User Registered // userId,email,password                     "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Event:
		properties := result.Lines[0].(emd.Event).Properties
		if len(properties) != 3 {
			t.Error("Unexpected number of Event.Properties")
		}
		if properties[0].Name != "userId" {
			t.Error("Event property not found")
		}
		if properties[1].Name != "email" {
			t.Error("Event property not found")
		}
		if properties[2].Name != "password" {
			t.Error("Event property not found")
		}
	default:
		t.Error("expected event")
	}
}

func TestShouldGetPropertiesWithTrailingComma(t *testing.T) {
	input := []string{"User Registered // userId,email,password,                     "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Event:
		properties := result.Lines[0].(emd.Event).Properties
		if len(properties) != 3 {
			t.Error("Unexpected number of Event.Properties")
		}
		if properties[0].Name != "userId" {
			t.Error("Event property not found")
		}
		if properties[1].Name != "email" {
			t.Error("Event property not found")
		}
		if properties[2].Name != "password" {
			t.Error("Event property not found")
		}
	default:
		t.Error("expected event")
	}
}

func TestShouldGetEventWithoutProperties(t *testing.T) {
	input := []string{"User Registered                    "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Event:
		if result.Lines[0].(emd.Event).Name != "User Registered" {
			t.Error("Unexpected Event.Name")
		}
	default:
		t.Error("expected event")
	}
}

func TestShouldGetEventWithoutPropertiesWithTrailingSlashes(t *testing.T) {
	input := []string{"User Registered //  "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Event:
		if result.Lines[0].(emd.Event).Name != "User Registered" {
			t.Error("Unexpected Event.Name")
		}
	default:
		t.Error("expected event")
	}
}

func TestShouldNotReturnEventWhenCommandGiven(t *testing.T) {
	input := []string{"Command This->"}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) > 0 {
		switch result.Lines[0].(type) {
		case emd.Event:
			t.Error("Unexpected Event.")
		}
	}
}
func TestShouldNotReturnEventWhenDocumentGiven(t *testing.T) {
	input := []string{"A Read Model* // one,two"}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) > 0 {
		switch result.Lines[0].(type) {
		case emd.Event:
			t.Error("Unexpected Event.")
		}
	}
}
