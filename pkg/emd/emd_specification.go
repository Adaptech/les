package emd

// Emd is a specification written in Event Markdown.
type Emd struct {
	Lines  []Item
	Errors []EmdValidationError `yaml:"Errors"` // The presence of errors means that no conversion to

}

// Item is a line item in an Event Markdown file.
type Item interface {
}

// A Comment in Event Markdown
// Example:
// "Order Placed"
type Comment struct {
	Text string
}

// Command in Event Markdown language
// Example:
// "Place Order-> // orderId, placedDate, deliveryDate"
type Command struct {
	Name       string
	Parameters []Parameter
}

// A Parameter for an emd command
type Parameter struct {
	Name string
}

// Event describes an emd event
type Event struct {
	Name       string
	Properties []Property
}

// Document describes an emd read model document
type Document struct {
	Name       string
	Properties []Property
}

// A Property of an emd event
type Property struct {
	Name string
}

// A EmdValidationError means that the event markdown structure cannot be used to generate event markup.
type EmdValidationError struct {
	ErrorID string
	Message string
}
