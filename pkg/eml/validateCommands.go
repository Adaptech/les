package eml

import (
	"fmt"
	"strings"
	"unicode"
)

func (c *Solution) validateCommands(boundedContextName string, stream Stream, events map[string]Event, readmodels map[string]Readmodel) {
	if len(stream.Commands) == 0 {
		validationError := ValidationError{
			ErrorID: "NoCommands",
			Context: boundedContextName,
			Stream:  stream.Name,
			Message: "Event stream '" + stream.Name + "': Must have at least one command which publishes event(s) to it.",
		}
		c.Errors = append(c.Errors, validationError)
	}
	for _, command := range stream.Commands {
		if len(strings.Trim(command.Command.Name, " ")) == 0 {
			validationError := ValidationError{
				ErrorID: "NoCommandName",
				Context: boundedContextName,
				Stream:  stream.Name,
				Message: "Command Name cannot be blank: " + command.Command.Name,
			}
			c.Errors = append(c.Errors, validationError)
		}

		if !isValidCommandName(command.Command.Name) {
			validationError := ValidationError{
				ErrorID: "InvalidCommandName",
				Context: boundedContextName,
				Stream:  stream.Name,
				Message: "Invalid command name '" + command.Command.Name + "': Names must start with an alpha character and have no leading or trailing spaces.",
			}
			c.Errors = append(c.Errors, validationError)
		}

		for _, parameter := range command.Command.Parameters {
			// Verify that parameters have parameter types:
			if len(strings.Trim(parameter.Type, " ")) == 0 {
				propertyTypeValidationError := ValidationError{
					ErrorID: "NoParameterType",
					Context: boundedContextName,
					Stream:  stream.Name,
					Message: "Command parameter types cannot be blank: " + command.Command.Name,
				}
				c.Errors = append(c.Errors, propertyTypeValidationError)
			}

			if !isValidName(parameter.Name) {
				parameterNameError := ValidationError{
					ErrorID: "InvalidCommandParameterName",
					Context: boundedContextName,
					Stream:  stream.Name,
					Message: "Invalid command parameter name '" + command.Command.Name + "': Names must start with an alpha character and contain no spaces.",
				}
				c.Errors = append(c.Errors, parameterNameError)
			}

			// if there is a MustExistIn <readmodel> rule, ensure that the read model exists
			if parameter.RuleExists("MustExistIn") {
				readModel := parameter.MustExistInReadmodel()
				if readModel != "" {
					readmodelFound, readmodelExists := readmodels[readModel]
					if !readmodelExists {
						propertyError := ValidationError{
							ErrorID: "MustExistInReadmodelNotFound",
							Context: boundedContextName,
							Stream:  stream.Name,
							Message: "'" + command.Command.Name + " " + parameter.Name + "' has a MustExistIn " + readModel + " rule, but that read model does not exist.",
						}
						c.Errors = append(c.Errors, propertyError)
					} else {
						aggregatePartOfParameterName := getStreamIDPartOf(parameter.Name)
						if readmodelFound.Readmodel.Key != aggregatePartOfParameterName {
							propertyError := ValidationError{
								ErrorID: "MustExistInReadmodelNotFound",
								Context: boundedContextName,
								Stream:  stream.Name,
								Message: "'" + command.Command.Name + " " + parameter.Name + "' has a MustExistIn " + readModel + " rule, but that read model does not exist.",
							}
							c.Errors = append(c.Errors, propertyError)
						}
					}
				} else {
					// MustExistInHasNoReadmodelArg
					propertyError := ValidationError{
						ErrorID: "MustExistInHasNoReadmodelArgument",
						Context: boundedContextName,
						Stream:  stream.Name,
						Message: "'" + command.Command.Name + " " + parameter.Name + "' has a MustExistIn " + readModel + " rule, but no readmodel name argument.",
					}
					c.Errors = append(c.Errors, propertyError)
				}
			}
		}
		if len(command.Command.Postconditions) == 0 {
			propertyTypeValidationError := ValidationError{
				ErrorID: "CommandHasNoEvents",
				Context: boundedContextName,
				Stream:  stream.Name,
				Message: "'" + command.Command.Name + "' does not produce any events. Commands must result in state changes.",
			}
			c.Errors = append(c.Errors, propertyTypeValidationError)
		}
		for _, postcondition := range command.Command.Postconditions {
			_, eventExists := events[postcondition]
			if !eventExists {
				propertyTypeValidationError := ValidationError{
					ErrorID: "PostconditionMustExist",
					Context: boundedContextName,
					Stream:  stream.Name,
					Message: command.Command.Name + ": Unknown postcondition: The command cannot result in an event of type '" + postcondition + "' because the event type doesn't exist.",
				}
				c.Errors = append(c.Errors, propertyTypeValidationError)
			}
		}

	}
}

func getStreamIDPartOf(parameterName string) string {
	positionOfFirstUppercaseLetter := firstUppercaseLetterIn(parameterName)
	if positionOfFirstUppercaseLetter == len(parameterName)-2 {
		return parameterName
	}
	streamIDPartOfParameterName := parameterName[positionOfFirstUppercaseLetter:]
	a := []rune(streamIDPartOfParameterName)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func firstUppercaseLetterIn(s string) int {
	for i, c := range s {
		char := fmt.Sprintf("%c", c)
		if char == strings.ToUpper(char) {
			return i
		}
	}
	return -1
}
