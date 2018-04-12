package emd_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/emd"
)

func TestShouldIgnoreLinesWithLeadingSlashes(t *testing.T) {
	input := []string{"// User Registered  "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) > 0 {
		t.Error("unexpected command, event or document found")
		return
	}
}

func TestShouldIgnoreLinesWithLessThan2Characters(t *testing.T) {
	input := []string{"xx"}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) > 0 {
		t.Error("unexpected command, event or document found")
		return
	}
}
func TestShouldSkipBlankLines(t *testing.T) {
	input := []string{""}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) > 0 {
		t.Error("unexpected command, event or document found")
		return
	}
}
