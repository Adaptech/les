package eml_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/eml"
)

func TestValidReadmodelHasNoValidationErrors(t *testing.T) {
	const emlYAML = `Solution: Identity & Access
Contexts:
- Name: Identity & Access
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: RegisterUser
        Parameters:
        - Name: userId
          Type: string
          IsRequired: true
        - Name: email
          Type: string
          IsRequired: false
        - Name: password
          Type: string
          IsRequired: false
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: userId
          Type: string
          IsHashed: false
        - Name: email
          Type: string
          IsHashed: false
        - Name: password
          Type: string
          IsHashed: true
  Readmodels:
  - Readmodel:
      Name: Users
      Key: userId
      SubscribesTo:
      - UserRegistered
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if len(sut.Errors) != 0 {
		t.Error("expected no errors")
	}
}

func TestMustHaveReadmodelKey(t *testing.T) {
	readmodel := eml.Readmodel{}
	readmodel.Readmodel.Name = "asdf"
	subscribesToEvents := []string{"InvoiceCreated"}
	readmodel.Readmodel.SubscribesTo = subscribesToEvents
	readmodels := []eml.Readmodel{}
	readmodels = append(readmodels, readmodel)
	context := eml.BoundedContext{Name: "Invoice", Readmodels: readmodels}
	contexts := []eml.BoundedContext{}
	contexts = append(contexts, context)

	system := eml.Solution{
		Name:     "Accounting Solution",
		Contexts: contexts,
	}
	system.Validate()
	if !hasError("MissingReadmodelKey", system.Errors) {
		t.Error("expected error")
	}
}

func Test_readmodel_key_must_exist(t *testing.T) {
	const emlYAML = `EmlVersion: 0.1-alpha
Solution: Hello World
Contexts:
- Name: Hello World
  Streams:
  - Stream: Hello
    Commands:
    - Command:
        Name: SayHello
        Parameters:
        - Name: text
          Type: string
          Rules: []
        - Name: helloId
          Type: string
          Rules:
          - IsRequired
        Postconditions:
        - HelloSaid
    Events:
    - Event:
        Name: HelloSaid
        Properties:
        - Name: text
          Type: string
        - Name: helloId
          Type: string
  Readmodels:
  - Readmodel:
      Name: ListOfHellos
      Key: propertyWhichDoesntExistId
      SubscribesTo:
      - HelloSaid
Errors: []
`
	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))
	sut.Validate()
	if !hasError("ReadModelKeyDoesNotExist", sut.Errors) {
		t.Error("expected error")
	}
}

func TestMustHaveReadmodelName(t *testing.T) {
	const emlYAML = `Solution: User Registration
Contexts:
- Name: User Registration
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: RegisterUseremail
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
        - Name: email
          Type: string
          IsHashed: false
        - Name: userId
          Type: string
          IsHashed: false
  Readmodels:
  - Readmodel:
      Key: userId
      SubscribesTo:
      - UserRegistered
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("MissingReadmodelName", sut.Errors) {
		t.Error("expected error")
	}
}

func TestReadmodelMustHaveSubscribesToEvent(t *testing.T) {
	readmodel := eml.Readmodel{}
	readmodel.Readmodel.Key = "InvoiceId"
	readmodel.Readmodel.Name = "SomeReadModel"
	readmodels := []eml.Readmodel{}
	readmodels = append(readmodels, readmodel)
	context := eml.BoundedContext{Name: "Invoice", Readmodels: readmodels}
	contexts := []eml.BoundedContext{}
	contexts = append(contexts, context)

	system := eml.Solution{
		Name:     "Accounting Solution",
		Contexts: contexts,
	}
	system.Validate()
	if !hasError("MissingReadmodelSubscribesToEvent", system.Errors) {
		t.Error("expected error")
	}
}

func TestReadmodelSubscribesToEventsMustExist(t *testing.T) {
	const emlYAML = `Solution: Identity & Access
Contexts:
- Name: Identity & Access
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: RegisterUser
        Parameters:
        - Name: userId
          Type: string
          IsRequired: true
        - Name: email
          Type: string
          IsRequired: false
        - Name: password
          Type: string
          IsRequired: false
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: userId
          Type: string
          IsHashed: false
        - Name: email
          Type: string
          IsHashed: false
        - Name: password
          Type: string
          IsHashed: true
  Readmodels:
  - Readmodel:
      Name: Users
      Key: userId
      SubscribesTo:
      - ANonExistingEvent
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if !hasError("SubscribesToEventNotFound", sut.Errors) {
		t.Error("expected error")
	}
}

func Test_readmodel_name_with_spaces_is_ok(t *testing.T) {
	const emlYAML = `Solution: Identity & Access
Contexts:
- Name: Identity & Access
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: Register User
        Parameters:
        - Name: userId
          Type: string
          IsRequired: true
        - Name: email
          Type: string
          IsRequired: false
        - Name: password
          Type: string
          IsRequired: false
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: userId
          Type: string
          IsHashed: false
        - Name: email
          Type: string
          IsHashed: false
        - Name: password
          Type: string
          IsHashed: true
  Readmodels:
  - Readmodel:
      Name: Users List
      Key: userId
      SubscribesTo:
      - UserRegistered
Errors: []
`

	sut := eml.Solution{}
	sut.LoadYAML([]byte(emlYAML))

	sut.Validate()
	if len(sut.Errors) != 0 {
		t.Error("expected no errors")
	}
}
