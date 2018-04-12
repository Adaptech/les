// Package eml implements the Event Markup Language specification
package eml

// A Solution describes an event sourced system
// It will contain bounded contexts and meta information about what environment to deploy it to.
// It can also contain references to other bounded contexts from the 'PlayStore' (Context Store? Microservice Store?
type Solution struct {
	EmlVersion string            `yaml:"EmlVersion"`
	Name       string            `yaml:"Solution"`
	Contexts   []BoundedContext  `yaml:"Contexts"`
	Errors     []ValidationError `yaml:"Errors"` // The presence of errors means that no API can be generated for the Solution.
}

// A BoundedContext is a context in which a ubiquitous language applies.
type BoundedContext struct {
	Name       string      `yaml:"Name"`
	Streams    []Stream    `yaml:"Streams"`
	Readmodels []Readmodel `yaml:"Readmodels"`
}

// A Stream is a stream of events representing a transactional scope.
// In Domain Driven Design, this is known as an "aggregate".
// In comp sci, it can be represented as a state machine.
type Stream struct {
	Name     string    `yaml:"Stream"`
	Commands []Command `yaml:"Commands"`
	Events   []Event   `yaml:"Events"`
}

// An Event represents a fact that occurred as a result of a state change.
type Event struct {
	Event struct {
		Name       string     `yaml:"Name"`
		Properties []Property `yaml:"Properties"`
	} `yaml:"Event"`
}

// A Property of an event.
type Property struct {
	Name     string `yaml:"Name"`
	Type     string `yaml:"Type"`
	IsHashed bool   `yaml:"IsHashed,omitempty"`
}

// A Command in a bounded context
type Command struct {
	Command struct {
		Name       string      `yaml:"Name"`
		Parameters []Parameter `yaml:"Parameters"`
		// Preconditions  []Event
		Postconditions []string `yaml:"Postconditions"`
	} `yaml:"Command"`
}

// A Parameter for a command.
type Parameter struct {
	Name  string   `yaml:"Name"`
	Type  string   `yaml:"Type"`
	Rules []string `yaml:"Rules"`
}

// A Readmodel which subscribes to events
type Readmodel struct {
	Readmodel struct {
		Name         string   `yaml:"Name"`
		Key          string   `yaml:"Key"`
		SubscribesTo []string `yaml:"SubscribesTo"`
	} `yaml:"Readmodel"`
}

// A ValidationError means that the eml structure cannot be used to generate an API.
type ValidationError struct {
	ErrorID   string
	Context   string
	Stream    string
	Command   string
	Event     string
	Readmodel string
	Message   string
}
