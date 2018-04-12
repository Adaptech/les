package eml_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/eml"
)

func TestMustHaveSolutionName(t *testing.T) {
	system := eml.Solution{
		Name: "",
		Contexts: []eml.BoundedContext{
			eml.BoundedContext{Name: "identity and access"},
		},
	}
	system.Validate()
	if !hasError("NoSolutionName", system.Errors) {
		t.Error("expected error")
	}
}
