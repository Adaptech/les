package eml

import (
	"strings"
)

type preconditionType struct {
	Type   string
	Tokens []string
}

func parsePrecondition(precondition string) *preconditionType {
	tokens := tokenize(precondition)
	preconditionRule := determinePreconditionTypeFrom(tokens)
	switch preconditionRule {
	case "MustHaveHappened":
		// Example: "UserRegistered MustHaveHappened"
		return &preconditionType{Type: "MustHaveHappened", Tokens: tokens}
	default:
		return nil
	}
}

func tokenize(text string) []string {
	return strings.Split(text, " ")
}

func determinePreconditionTypeFrom(tokens []string) string {
	if len(tokens) == 2 && tokens[1] == "MustHaveHappened" {
		// Example: "UserRegistered MustHaveHappened"
		return "MustHaveHappened"
	}
	return ""
}
