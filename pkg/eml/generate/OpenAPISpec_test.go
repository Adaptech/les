package generate_test

import (
	"strings"
	"testing"

	"github.com/Adaptech/les/pkg/eml"
	"github.com/Adaptech/les/pkg/eml/generate"
)

func TestMustHaveTitleAndDescription(t *testing.T) {
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
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: userId
          Type: string
          IsHashed: false
        - Name: name
          Type: string
          IsHashed: false
  Readmodels:
  - Readmodel:
      Name: Users
      Key: userId
      SubscribesTo:
      - UserRegistered
Errors: []
`
	const expectedResult = `openapi: "3.0.1"

info:
  title: "Identity & Access"
  description: "Identity & Access API"
  version: "0.1"

servers:
  - url: http://localhost:3001/api/v1
    description: localhost

paths:`

	solution := eml.Solution{}
	solution.LoadYAML([]byte(emlYAML))
	result := generate.OpenAPISpec(solution)
	if !strings.Contains(result, expectedResult) {
		t.Log(result)
		t.Error("openapi/info/servers/path sections not found.")
	}

}

func TestMustHaveCommand(t *testing.T) {
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
        Postconditions:
        - UserRegistered
    Events:
    - Event:
        Name: UserRegistered
        Properties:
        - Name: userId
          Type: string
          IsHashed: false
        - Name: name
          Type: string
          IsHashed: false
  Readmodels:
  - Readmodel:
      Name: Users
      Key: userId
      SubscribesTo:
      - UserRegistered
Errors: []
`
	const expectedResult = `    /User/RegisterUser:
      post:
        tags: 
          - User
        summary: RegisterUser
        description: RegisterUser command.
        requestBody:
          required: true
          content:
            application/json:
              schema:
                type: object
`

	solution := eml.Solution{}
	solution.LoadYAML([]byte(emlYAML))
	result := generate.OpenAPISpec(solution)
	if !strings.Contains(result, expectedResult) {
		t.Log(result)
		t.Error("command's paths.post OpenAPI section not found.")
	}

}

func TestMustHaveCommandParameters(t *testing.T) {
	const emlYAML = `Solution: Identity & Access
Contexts:
- Name: Identity & Access
  Streams:
  - Stream: User
    Commands:
    - Command:
        Name: RegisterUser
        Parameters:
        - Name: name
          Type: string
          IsRequired: false
        - Name: email
          Type: string
          IsRequired: true
        - Name: userId
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
        - Name: name
          Type: string
          IsHashed: false
  Readmodels:
  - Readmodel:
      Name: Users
      Key: userId
      SubscribesTo:
      - UserRegistered
Errors: []
`
	solution := eml.Solution{}
	solution.LoadYAML([]byte(emlYAML))
	result := generate.OpenAPISpec(solution)
	if !strings.Contains(result, "properties:") || !strings.Contains(result, "name:") || !strings.Contains(result, "email:") || !strings.Contains(result, "userId:") {
		t.Log(result)
		t.Error("command parameters not found.")
	}
	if !strings.Contains(result, "responses:") {
		t.Log(result)
		t.Error("command respohses.")
	}
}
