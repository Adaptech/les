package rules_test

import (
	"testing"

	"github.com/Adaptech/les/pkg/eml"
	"github.com/Adaptech/les/pkg/eml/rules"
)

func Test_should_recognize_valid_mustexistin_rule(t *testing.T) {
	sut := "MustExistIn UserLookup fromUserId"

	validationErrors, _ := rules.Validator(sut)
	if len(validationErrors) > 0 {
		t.Error("Expected no validation errors.")
	}
}

func Test_should_fail_missing_mustexist_readmodel_operand(t *testing.T) {
	sut := "MustExistIn UserLookup fromUserId"

	validationErrors, _ := rules.Validator(sut)
	if hasError("MissingParameterReadmodel", validationErrors) {
		t.Error("Expected MissingParameterReadmodel error.")
	}
}

func hasError(errorId string, errors []eml.ValidationError) bool {
	for _, err := range errors {
		if err.ErrorID == errorId {
			return true
		}
	}
	return false
}
