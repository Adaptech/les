package eml_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/eml"
)

// Preconditions:
// - Not TimesheetCreated
// - LastEventIs Not TimesheetSubmitted

func Test_command_precondition_event_must_belong_to_the_same_stream_as_the_command(t *testing.T) {
	const emlYAML = `Solution: User Registration
Contexts:
- Name: User Registration
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: RegisterUser
        Parameters:
        - Name: email
          Type: string
          IsRequired: true
        Preconditions:
        - SomeEventThatDoesntExist MustHaveHappened
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: "name"
          Type: string
          IsHashed: false
        - Name: "userId"
          Type: string
          IsHashed: false
  Readmodels: []
Errors: []
`
	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))
	sut.Validate()
	if !hasError("PreconditionEventDoesntExistInStream", sut.Errors) {
		t.Error("expected error")
	}
}

func Test_command_precondition_event_which_belonts_to_the_same_stream_as_the_command_passes_validation(t *testing.T) {
	const emlYAML = `Solution: User Registration
Contexts:
- Name: User Registration
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: RegisterUser
        Parameters:
        - Name: email
          Type: string
          IsRequired: true
        Postconditions:
        - UserRegistered
    - Command:
        Name: Authenticate
        Parameters:
        - Name: email
          Type: string
          IsRequired: true
        Preconditions:
        - UserRegistered MustHaveHappened
        Postconditions:
        - UserAuthenticated
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: "name"
          Type: string
          IsHashed: false
        - Name: "userId"
          Type: string
          IsHashed: false
    - Event:
        Name: UserAuthenticated
        Properties:
        - Name: "userId"
          Type: string
          IsHashed: false
  Readmodels: []
Errors: []
`
	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))
	sut.Validate()
	if len(sut.Errors) > 0 {
		t.Error("unexpected error")
	}
}

func Test_unknown_command_precondition_type_results_in_validation_error(t *testing.T) {
	const emlYAML = `Solution: User Registration
Contexts:
- Name: User Registration
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: RegisterUser
        Parameters:
        - Name: email
          Type: string
          IsRequired: true
        Preconditions:
        - asdfsadfasdfasdfasdf
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: "name"
          Type: string
          IsHashed: false
        - Name: "userId"
          Type: string
          IsHashed: false
  Readmodels: []
Errors: []
`
	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))
	sut.Validate()
	if !hasError("UnknownCommandPreconditionType", sut.Errors) {
		t.Error("expected error")
	}
}
