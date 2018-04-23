package eml_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/eml"
)

func TestMustHaveEvents(t *testing.T) {
	context := eml.BoundedContext{Name: "User Registration"}
	stream := eml.Stream{Name: "User"}
	context.Streams = append(context.Streams, stream)

	contexts := []eml.BoundedContext{}
	contexts = append(contexts, context)

	system := eml.Solution{
		Name:     "Accounting System",
		Contexts: contexts,
	}
	system.Validate()
	if !hasError("NoEvents", system.Errors) {
		t.Error("expected error")
	}
}

func TestEventMustHaveAggregateId(t *testing.T) {

	const emlYAML = `Solution: User Registration
Contexts:
- Name: User Registration
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: 'RegisterUser'
        Parameters:
        - Name: UserId
          Type: string
          IsRequired: true
        Postconditions:
        - UserRegistered
        - UserAuthenticated
    Events:
    - Event:
        Name: UserRegistered
        Properties: []
    - Event:
        Name: UserAuthenticated
        Properties:
        - Name: userId
          Type: string
          IsHashed: false
  Readmodels: []
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("NoAggregateId", sut.Errors) {
		t.Error("expected error")
	}
	if len(sut.Errors) != 1 {
		t.Error("expected different number of errors")
	}
}

func Test_event_name_should_start_with_alpha(t *testing.T) {
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
        Name: 124Vegetables Stocked
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
	if !hasError("InvalidEventName", sut.Errors) {
		t.Error("expected error")
	}
}

func TestPropertyNameMustNotBeBlank(t *testing.T) {

	const emlYAML = `Solution: User Registration
Contexts:
- Name: User Registration
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: RegisterUser
        Parameters:
        - Name: userId
          Type: string
          IsRequired: true
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: "   "
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
	if !hasError("NoPropertyName", sut.Errors) {
		t.Error("expected error")
	}
	if !hasError("InvalidEventPropertyName", sut.Errors) {
		t.Error("expected error")
	}
	if len(sut.Errors) != 2 {
		t.Error("expected different number of errors")
	}
}
func TestPropertyTypesMustNotBeBlank(t *testing.T) {

	const emlYAML = `Solution: User Registration
Contexts:
- Name: User Registration
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: RegisterUser
        Parameters:
        - Name: userId
          Type: string
          IsRequired: true
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: "name"
          Type: 
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
	if !hasError("NoPropertyType", sut.Errors) {
		t.Error("expected error")
	}
	if len(sut.Errors) != 1 {
		t.Error("expected different number of errors")
	}
}

func TestPropertiesShouldStartWithAlpha(t *testing.T) {
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
        - Name: 2344vegetablesId
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
	if !hasError("InvalidEventPropertyName", sut.Errors) {
		t.Error("expected error")
	}
}

func TestEventNamesMustNotBeBlank(t *testing.T) {
	const emlYAML = `Solution: Veggies Galore
Contexts:
- Name: Veggies Galore
  Streams:
  - Stream: Vegetables
    Commands:
    - Command:
        Name: 12345
        Parameters:
        - Name: VegetablesId
          Type: string
          IsRequired: true
        Postconditions:
        - VegetablesStocked/
    Events:
    - Event:
        Name: 
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
	if !hasError("NoEventID", sut.Errors) {
		t.Error("expected error")
	}
}

func Test_event_must_be_a_result_of_executing_command(t *testing.T) {

	const emlYAML = `Solution: User Registration
Contexts:
- Name: User Registration
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: 'RegisterUser'
        Parameters:
        - Name: UserId
          Type: string
          IsRequired: true
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: userId
          Type: string
          IsHashed: false
    - Event:
        Name: UserAuthenticated
        Properties:
        - Name: userId
          Type: string
          IsHashed: false
  Readmodels: []
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("EventMustBeCommandPostcondition", sut.Errors) {
		t.Error("expected error")
	}
	if len(sut.Errors) != 1 {
		t.Error("expected different number of errors")
	}
}
