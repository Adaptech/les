package emd_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/emd"
)

func TestShouldGetCommandWithParameters(t *testing.T) {
	input := []string{"  Validate Registration  -> // userId                     "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no command found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Command:
		if result.Lines[0].(emd.Command).Name != "Validate Registration" {
			t.Error("Unexpected Command.Name")
		}
	default:
		t.Error("expected command")
	}
}

func TestShouldFindCommandParameters(t *testing.T) {
	input := []string{"Validate->// userId                    "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no command found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Command:
		parameters := result.Lines[0].(emd.Command).Parameters
		if len(parameters) != 1 {
			t.Error("Unexpected number of command.Parameters")
		}
		if parameters[0].Name != "userId" {
			t.Error("Command parameter not found")
		}
	default:
		t.Error("expected command")
	}
}

func TestShouldGetCommandWithoutProperties(t *testing.T) {
	input := []string{"Register->                   "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no event found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Command:
		if result.Lines[0].(emd.Command).Name != "Register" {
			t.Error("Unexpected Command.Name")
		}
	default:
		t.Error("expected command")
	}
}

func TestShouldGetCommandWithoutPropertiesWithTrailingSlashes(t *testing.T) {
	input := []string{"Register->//  "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no command found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Command:
		if result.Lines[0].(emd.Command).Name != "Register" {
			t.Error("Unexpected Command.Name")
		}
	default:
		t.Error("expected command")
	}
}
