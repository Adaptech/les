package eml_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/eml"
)

func TestMustHaveAggregateEventStreams(t *testing.T) {
	context := eml.BoundedContext{Name: "User Registration"}
	contexts := []eml.BoundedContext{}
	contexts = append(contexts, context)

	system := eml.Solution{
		Name:     "Accounting System",
		Contexts: contexts,
	}
	system.Validate()
	if !hasError("NoStreams", system.Errors) {
		t.Error("expected error")
	}
}
