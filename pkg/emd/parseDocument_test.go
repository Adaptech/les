package emd_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/emd"
)

func TestShouldGetDocumentWithParameters(t *testing.T) {
	input := []string{"  Validate Registration  -> // userId                     "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no command found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Command:
		if result.Lines[0].(emd.Command).Name != "Validate Registration" {
			t.Error("Unexpected Command.Name")
		}
	default:
		t.Error("expected command")
	}
}

func TestShouldGetDocumentParameters(t *testing.T) {
	input := []string{"Users* // userId,email"}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no document found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Document:
		properties := result.Lines[0].(emd.Document).Properties
		if len(properties) != 2 {
			t.Error("Unexpected number of document.Properties")
		}
		if properties[0].Name != "userId" {
			t.Error("Document property not found")
		}
		if properties[1].Name != "email" {
			t.Error("Document property not found")
		}
	default:
		t.Error("expected document")
	}
}

func TestShouldGetDocumentParametersWithTrailingComma(t *testing.T) {
	input := []string{"Users* // userId,email  ,     "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no document found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Document:
		properties := result.Lines[0].(emd.Document).Properties
		if len(properties) != 2 {
			t.Error("Unexpected number of document.Properties")
		}
		if properties[0].Name != "userId" {
			t.Error("Document property not found")
		}
		if properties[1].Name != "email" {
			t.Error("Document property not found")
		}
	default:
		t.Error("expected document")
	}
}

func TestShouldGetDocumentProperties(t *testing.T) {
	input := []string{"Register* // userId,email,password                     "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no document found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Document:
		properties := result.Lines[0].(emd.Document).Properties
		if len(properties) != 3 {
			t.Error("Unexpected number of Document.Properties")
		}
		if properties[0].Name != "userId" {
			t.Error("Document property not found")
		}
		if properties[1].Name != "email" {
			t.Error("Document property not found")
		}
		if properties[2].Name != "password" {
			t.Error("Document property not found")
		}
	default:
		t.Error("expected document")
	}
}

func TestShouldGetDocumentWithoutProperties(t *testing.T) {
	input := []string{"User Register*                    "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no document found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Document:
		if result.Lines[0].(emd.Document).Name != "User Register" {
			t.Error("Unexpected Document.Name")
		}
	default:
		t.Error("expected document")
	}
}

func TestShouldGetDocumentWithoutPropertiesWithTrailingSlashes(t *testing.T) {
	input := []string{"A List* //"}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}
	if len(result.Lines) == 0 {
		t.Error("no document found")
		return
	}
	switch result.Lines[0].(type) {
	case emd.Document:
		if result.Lines[0].(emd.Document).Name != "A List" {
			t.Error("Unexpected Document.Name")
		}
	default:
		t.Error("expected document")
	}
}
