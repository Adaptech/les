package eml

import (
	"strings"
)

// RuleExists verifies whether a command parameter validation rule exists.
func (r *Parameter) RuleExists(rule string) bool {
	for _, existingRule := range r.Rules {
		if strings.HasPrefix(string(existingRule), rule) {
			return true
		}
	}
	return false
}

// MustExistInReadmodel gets a rule and it's parameters if it exists, otherwise nil.
func (r *Parameter) MustExistInReadmodel() string {
	for _, existingRule := range r.Rules {
		if strings.HasPrefix(string(existingRule), "MustExistIn") {
			ruleParts := strings.Split(string(existingRule), " ")
			if len(ruleParts) >= 2 {
				return ruleParts[1]
			}
			return ""
		}
	}
	return ""
}
