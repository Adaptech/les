package emd_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/emd"
)

func TestComments(t *testing.T) {
	input := []string{"# This is a comment                     "}
	result, err := emd.Parse(input)
	if err != nil {
		panic(err)
	}

	switch result.Lines[0].(type) {
	case emd.Comment:
		if result.Lines[0].(emd.Comment).Text != "This is a comment" {
			t.Error("Unexpected Comment.Text")
		}
	default:
		t.Error("expected comment")

	}
}
