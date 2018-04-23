package eml_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/eml"
)

func TestMustHaveCommands(t *testing.T) {
	context := eml.BoundedContext{Name: "User Registration"}
	stream := eml.Stream{Name: "User"}
	event := eml.Event{}
	event.Event.Name = "123"
	event.Event.Properties = []eml.Property{eml.Property{Name: "userId", Type: "string"}}
	stream.Events = append(stream.Events, event)

	context.Streams = append(context.Streams, stream)

	contexts := []eml.BoundedContext{}
	contexts = append(contexts, context)

	system := eml.Solution{
		Name:     "Accounting System",
		Contexts: contexts,
	}
	system.Validate()
	if !hasError("NoCommands", system.Errors) {
		t.Error("expected error")
	}
}

func TestParameterTypesMustNotBeBlank(t *testing.T) {
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
          IsRequired: true
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
	if !hasError("NoParameterType", sut.Errors) {
		t.Error("expected error")
	}
	if len(sut.Errors) != 1 {
		t.Error("expected different number of errors")
	}
}

func Test_command_parameter_should_start_with_alpha(t *testing.T) {
	const emlYAML = `Solution: Veggies Galore
Contexts:
- Name: Veggies Galore
  Streams:
  - Stream: Vegetables
    Commands:
    - Command:
        Name: 'Stock Vegetable '
        Parameters:
        - Name: 12345VegetablesId
          Type: string
          IsRequired: true
        Postconditions:
        - 'Vegetables Stocked'
    Events:
    - Event:
        Name: Vegetables Stocked
        Properties:
        - Name: vegetablesId
          Type: string
          IsHashed: false
        - Name: name
          Type: string
          IsHashed: false
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("InvalidCommandParameterName", sut.Errors) {
		t.Error("expected error")
	}
}

func Test_command_parameter_should_not_contain_spaces(t *testing.T) {
	const emlYAML = `Solution: Veggies Galore
Contexts:
- Name: Veggies Galore
  Streams:
  - Stream: Vegetables
    Commands:
    - Command:
        Name: Stock Vegetable
        Parameters:
        - Name: Vegetables Id
          Type: string
          IsRequired: true
        Postconditions:
        - 'Vegetables Stocked'
    Events:
    - Event:
        Name: Vegetables Stocked
        Properties:
        - Name: vegetablesId
          Type: string
          IsHashed: false
        - Name: name
          Type: string
          IsHashed: false
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("InvalidCommandParameterName", sut.Errors) {
		t.Error("expected error")
	}

	if len(sut.Errors) != 1 {
		t.Error("expected only one error")
	}
}

func TestCommandNamesShouldStartWithAlpha(t *testing.T) {
	const emlYAML = `Solution: Veggies Galore
Contexts:
- Name: Veggies Galore
  Streams:
  - Stream: Vegetables
    Commands:
    - Command:
        Name: '$#@Stock Vegetable '
        Parameters:
        - Name: VegetablesId
          Type: string
          IsRequired: true
        Postconditions:
        - 'Vegetables Stocked'
    Events:
    - Event:
        Name: Vegetables Stocked
        Properties:
        - Name: vegetablesId
          Type: string
          IsHashed: false
        - Name: name
          Type: string
          IsHashed: false
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("InvalidCommandName", sut.Errors) {
		t.Error("expected error")
	}
}

func Test_command_names_can_contain_spaces(t *testing.T) {
	const emlYAML = `Solution: Veggies Galore
Contexts:
- Name: Veggies Galore
  Streams:
  - Stream: Vegetables
    Commands:
    - Command:
        Name: 'Stock Vegetable '
        Parameters:
        - Name: VegetablesId
          Type: string
          IsRequired: true
        Postconditions:
        - 'Vegetables Stocked'
    Events:
    - Event:
        Name: Vegetables Stocked
        Properties:
        - Name: vegetablesId
          Type: string
          IsHashed: false
        - Name: name
          Type: string
          IsHashed: false
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if len(sut.Errors) > 0 {
		t.Error("expected no errors")
		t.Log(sut.Errors)
	}
}

func TestCommandNameMustNotBeBlank(t *testing.T) {

	const emlYAML = `Solution: Veggies Galore
Contexts:
- Name: Veggies Galore
  Streams:
  - Stream: Vegetables
    Commands:
    - Command:
        Name:
        Parameters:
        - Name: VegetablesId
          Type: string
          IsRequired: true
        Postconditions:
        - VegetablesStocked/
    Events:
    - Event:
        Name: "VegetablesStocked"
        Properties:
        - Name: vegetablesId
          Type: string
          IsHashed: false
        - Name: name
          Type: string
          IsHashed: false
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("NoCommandName", sut.Errors) {
		t.Error("expected error")
	}
}

func TestPostconditionEventsMustExist(t *testing.T) {

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
        - UnknownEvent
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
	if !hasError("PostconditionMustExist", sut.Errors) {
		t.Error("expected error")
	}
	if len(sut.Errors) != 1 {
		t.Error("expected different number of errors")
	}
}

func Test_existing_postcondition_events_dont_result_in_validation_error(t *testing.T) {
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
	if len(sut.Errors) != 0 {
		t.Error("expected no errors")
	}
}

func TestCommandsMustResultInEvents(t *testing.T) {

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
    - Command:
        Name: DeleteUser
        Parameters:
        - Name: email
          Type: string
          IsRequired: true
        Postconditions:
        - UserDeleted
    Events:
    - Event:
        Name: UserDeleted
        Properties:
        - Name: "userId"
          Type: string
  Readmodels: []
Errors: []
`
	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("CommandHasNoEvents", sut.Errors) {
		t.Error("expected error")
	}
	if len(sut.Errors) != 1 {
		t.Error("expected different number of errors")
	}
}

func Test_MustExistIn_rule_has_readmodel_argument(t *testing.T) {
	const emlYAML = `Solution: Timesheets & Billing
Contexts:
- Name: Timesheets & Billing
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: RegisterUser
        Parameters:
        - Name: email
          Type: string
          Rules: []
        - Name: password
          Type: string
          Rules: []
        - Name: userId
          Type: string
          Rules:
          - IsRequired
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: email
          Type: string
        - Name: password
          Type: string
          IsHashed: true
        - Name: userId
          Type: string
  - Stream: Timesheet
    Commands:
    - Command:
        Name: CreateTimesheet
        Parameters:
        - Name: userId
          Type: string
          Rules:
          - MustExistIn 
        - Name: description
          Type: string
          Rules: []
        - Name: timesheetId
          Type: string
          Rules:
          - IsRequired
        Postconditions:
        - TimesheetCreated
    Events:
    - Event:
        Name: TimesheetCreated
        Properties:
        - Name: userId
          Type: string
        - Name: description
          Type: string
        - Name: timesheetId
          Type: string
  Readmodels:
  - Readmodel:
      Name: UserLookup
      Key: userId
      SubscribesTo:
      - UserRegistered
Errors: []
`
	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("MustExistInHasNoReadmodelArgument", sut.Errors) {
		t.Error("expected MustExistInHasNoReadmodelArg error")
	}
	if len(sut.Errors) != 1 {
		t.Error("expected different number of errors")
	}
}
