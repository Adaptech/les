package rules

import "github.com/Adaptech/les/pkg/eml"

// Validator determines if an input contains valid rules.
func Validator(input string) ([]eml.ValidationError, error) {
	return nil, nil
}

// for each token in the reversed prefix expression:
//   if token is an operator:
//     operand_1 ← pop from the stack
//     operand_2 ← pop from the stack
//     result ← evaluate token with operand_1 and operand_2
//     push result back onto the stack
//   else if token is an operand:
//     push token onto the stack
// result ← pop from the stack
