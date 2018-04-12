package eml_test

import (
	"strings"
	"testing"

	"github.com/Adaptech/les/pkg/eml"
)

func Test_must_have_boundedContext(t *testing.T) {
	system := eml.Solution{
		Name:     "Accounting Solution",
		Contexts: []eml.BoundedContext{}}
	system.Validate()
	if !hasError("NoBoundedContext", system.Errors) {
		t.Error("expected error")
	}
}

func Test_names_for_streams_commands_events_and_readmodels_must_be_unique_within_bounded_context(t *testing.T) {

	const emlYAML = `Solution: User Registration
Contexts:
- Name: User Registration
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: User
        Parameters:
        - Name: email
          IsRequired: true
          Type: string
        Postconditions:
        - User
    Events:
    - Event:
        Name: User
        Properties:
        - Name: "name"
          Type: string
          IsHashed: false
        - Name: "userId"
          Type: string
          IsHashed: false
  Readmodels: 
  - Readmodel:
      Name: User
      Key: userId
      SubscribesTo:
      - User  
Errors: []
`
	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("NamesMustBeUnique", sut.Errors) {
		t.Error("expected error")
	}
	if !strings.Contains(sut.Errors[0].Message, "4 times") {
		t.Error("Expected different error message.")
		t.Log(sut.Errors[0].Message)
	}
	if len(sut.Errors) != 1 {
		t.Error("expected different number of errors")
	}
}
