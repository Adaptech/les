package convert_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/convert"
	"github.com/Adaptech/les/pkg/emd"
	"github.com/Adaptech/les/pkg/eml"
)

func TestShouldCreateCommand(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register-> // userId, email, password",
		"User Registered",
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
	if len(streams[0].Commands) != 1 {
		t.Error("expected different no of User commands.")
	}
}

func TestCreateCommandParametersCorrect(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register-> // userId, email, password",
		"User Registered",
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
	userID := streams[0].Commands[0].Command.Parameters[0].Name
	if userID != "userId" {
		t.Error("Expected command name to be 'userId'). Found", userID)
	}
}

func TestShouldCreateOneCommandFor2EventPostconditions(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register-> // userId, email, password",
		"User Registered",
		"User Registration Confirmed",
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
	if len(streams[0].Commands) != 1 {
		t.Error("expected different no of User commands.")
	}
}
func TestCommandShouldHaveParameters(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register-> // userId, password",
		"User Registered // userId, password",
		"User Registration Confirmed // userId",
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
	if len(streams[0].Commands) != 1 {
		t.Error("expected different no of User commands.")
	} else {
		command := streams[0].Commands[0]
		if len(command.Command.Parameters) != 2 {
			t.Error("expected different no of User Register command parameters")
		} else {
			parameters := command.Command.Parameters
			if parameters[0].Name != "userId" {
				t.Error("expected different parameter name")
			}
			if parameters[0].Type != "string" {
				t.Error("expected different parameter type")
			}
			if parameters[1].Name != "password" {
				t.Error("expected different parameter name")
			}
			if parameters[1].Type != "string" {
				t.Error("expected different parameter type")
			}
		}
	}
}

func TestCommandShouldHavePostconditions(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register-> // userId, password",
		"User Registered",
		"User Registration Confirmed",
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
	if len(streams[0].Commands) != 1 {
		t.Error("expected different no of User commands.")
	} else {
		command := streams[0].Commands[0]
		if len(command.Command.Postconditions) != 2 {
			t.Error("expected different no of User Register command postconditions")
		}
	}
}

func TestCommandMustHaveAggregateIdBasedOnStreamName(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register->",
		"User Registered",
		"User Registration Confirmed",
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
	if len(streams[0].Commands) != 1 {
		t.Error("expected different no of User commands.")
	} else {
		command := streams[0].Commands[0]
		if len(command.Command.Parameters) != 1 {
			t.Error("expected one parameter: userId (= aggregate ID based on stream name")
		} else {
			parameter := command.Command.Parameters[0]
			if stringInSlice("IsRequired", parameter.Rules) == false {
				t.Error("expected aggregate ID to be a required field")
			}
			if parameter.Name != "userId" {
				t.Error("expected aggregate ID name to be stream name + Id (e.g. 'userId')")
			}
			if parameter.Type != "string" {
				t.Error("expected different aggregate ID type")
			}
		}
	}
}
func TestCommandMustHaveMustExistRuleForAggregateIdsOtherThanTheOneOfStreamTheCommandBelongsTo(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register-> // email",
		"User Registered // email",
		"Create Timesheet-> // timesheetId, userId, date",
		"Timesheet Created // timesheetId, userId, date",
		"UserLookup* // userId, date",
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
	if len(streams[0].Commands) != 1 {
		t.Error("expected different no of User commands.")
	} else {
		command := streams[1].Commands[0]
		if len(command.Command.Parameters) != 3 {
			t.Error("expected three parameters")
		} else {
			parameter := command.Command.Parameters[1]
			if len(parameter.Rules) != 1 {
				t.Error("expected only one rule")
			}
			if stringInSlice("MustExistIn UserLookup", parameter.Rules) == false {
				t.Error("expected MustExistIn UserLookup rule.")
			}
		}
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func TestCommandShouldNotGenerateAggregateIdIfOneAlreadyExists(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register-> // userId",
		"User Registered",
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
	if len(streams[0].Commands) != 1 {
		t.Error("expected different no of User commands.")
	} else {
		command := streams[0].Commands[0]
		if len(command.Command.Parameters) != 1 {
			t.Error("expected one parameter: userId (= aggregate ID based on stream name")
		} else {
			parameter := command.Command.Parameters[0]
			if stringInSlice("IsRequired", parameter.Rules) == false {
				t.Error("expected aggregate ID to be a required field")
			}
			if parameter.Name != "userId" {
				t.Error("expected aggregate ID name to be stream name + Id (e.g. 'userId')")
			}
			if parameter.Type != "string" {
				t.Error("expected different aggregate ID type")
			}
		}
	}
}

func TestAddingAggregateIdToCommandParametersShouldPreserveOtherParameters(t *testing.T) {
	input := []string{
		"# Timesheets",
		"Register-> // name",
		"User Registered",
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
	if len(streams[0].Commands) != 1 {
		t.Error("expected different no of User commands.")
	} else {
		command := streams[0].Commands[0]
		if len(command.Command.Parameters) != 2 {
			t.Error("expected different no. of parameters")
		} else {
			parameter0 := command.Command.Parameters[0]
			if parameter0.Name != "name" {
				t.Error("expected different parameter name")
			}
			parameter1 := command.Command.Parameters[1]
			if parameter1.Name != "userId" {
				t.Error("expected aggregateId (= userId) parameter")
			}
		}
	}
}

func TestParameterMustExistInFailsIfReadmodelDoesntExist(t *testing.T) {
	input := `Solution: Timesheets
Contexts:
- Name: Timesheets
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: Register
        Parameters:
        - Name: email
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
          IsHashed: false
        - Name: userId
          Type: string
          IsHashed: false
  - Stream: Timesheet
    Commands:
    - Command:
        Name: CreateTimesheet
        Parameters:
        - Name: timesheetId
          Type: string
          Rules:
          - IsRequired
        - Name: userId
          Type: string
          Rules:
          - MustExistIn UserLookup
        - Name: date
          Type: string
          Rules: []
        Postconditions:
        - TimesheetCreated
    Events:
    - Event:
        Name: TimesheetCreated
        Properties:
        - Name: timesheetId
          Type: string
          IsHashed: false
        - Name: userId
          Type: string
          IsHashed: false
        - Name: date
          Type: string
          IsHashed: false
  Readmodels:
    - Readmodel:
      Name: TimesheetHoursLookup
      Key: timesheethoursId
      SubscribesTo:
      - TimesheetCreated

`
	sut := eml.Solution{}
	sut.LoadYAML([]byte(input))
	sut.Validate()
	if !hasError("MustExistInReadmodelNotFound", sut.Errors) {
		t.Error("expected error")
	}
}

func TestParameterMustExistInReadmodelThatExistsButHasWrongKey(t *testing.T) {
	input := `Solution: Timesheets & Billing
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
          IsHashed: false
        - Name: password
          Type: string
          IsHashed: true
        - Name: userId
          Type: string
          IsHashed: false
  - Stream: Timesheet
    Commands:
    - Command:
        Name: CreateTimesheet
        Parameters:
        - Name: userId
          Type: string
          Rules:
          - MustExistIn UserLookup
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
          IsHashed: false
        - Name: description
          Type: string
          IsHashed: false
        - Name: timesheetId
          Type: string
          IsHashed: false
  Readmodels:
  - Readmodel:
      Name: UserLookup
      Key: timesheetId
      SubscribesTo:
      - UserRegistered
Errors: []
`
	sut := eml.Solution{}
	sut.LoadYAML([]byte(input))
	sut.Validate()
	if len(sut.Errors) != 1 {
		t.Error("Expected one error.")
	}
}

func TestParameterMustExistInReadmodelThatExistsAndHasKeySucceeds(t *testing.T) {
	input := `Solution: Timesheets & Billing
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
          IsHashed: false
        - Name: password
          Type: string
          IsHashed: true
        - Name: userId
          Type: string
          IsHashed: false
  - Stream: Timesheet
    Commands:
    - Command:
        Name: CreateTimesheet
        Parameters:
        - Name: userId
          Type: string
          Rules:
          - MustExistIn UserLookup
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
          IsHashed: false
        - Name: description
          Type: string
          IsHashed: false
        - Name: timesheetId
          Type: string
          IsHashed: false
  Readmodels:
  - Readmodel:
      Name: UserLookup
      Key: userId
      SubscribesTo:
      - UserRegistered
Errors: []
`
	sut := eml.Solution{}
	sut.LoadYAML([]byte(input))
	sut.Validate()
	if len(sut.Errors) > 0 {
		t.Error("Expected no errors.")
	}
}
func TestParameterWithAggregatePostfixMustExistInReadmodelThatExistsAndHasKeySucceeds(t *testing.T) {
	input := []string{
		"# Email System",
		"Register -> // userId, email, password",
		"User Registered // userId, email, hashedPassword",
		"Send Mail -> // emailId, fromUserId, toUserId, subject, body",
		"Email Sent // emailId, fromUserId, toUserId, subject, body",
		"UserLookup* // userId, email",
	}

	markdown, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	conversionResult, err := convert.EmdToEml(markdown)
	if err != nil {
		panic(err)
	}

	sut := conversionResult.Eml
	sut.Validate()
	if len(sut.Errors) > 0 {
		t.Error("Expected no errors.")
		t.Log(sut.Errors)
	}
}
