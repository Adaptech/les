package main

import (
	"fmt"

	"github.com/Adaptech/les/pkg/eml"
)

func checkIfFileContainsValidEml(inputFile string) (bool, error) {
	fileContent, err := ReadFile(inputFile)
	if err != nil {
		return false, err
	}
	markup := eml.Solution{}
	markup.LoadYAML(fileContent)
	markup.Validate()
	isValidEml := len(markup.Errors) == 0
	if !isValidEml {
		fmt.Println("EML Errors:")
		for _, validationError := range markup.Errors {
			printEmlError(validationError)
		}
	}
	return isValidEml, nil
}

func printEmlError(validationError eml.ValidationError) {
	var context string
	if validationError.Context != "" {
		context = validationError.Context
	}

	var stream string
	if validationError.Stream != "" {
		stream = "-" + validationError.Stream
	}

	var command string
	if validationError.Command != "" {
		command = "-" + validationError.Command
	}

	var event string
	if validationError.Event != "" {
		event = "-" + validationError.Event
	}

	var readModel string
	if validationError.Readmodel != "" {
		readModel = "-" + validationError.Readmodel
	}

	errorMessage := fmt.Sprintf("%s%s%s%s%s: %s\n",
		context,
		stream,
		command,
		event,
		readModel,
		validationError.Message)

	printError(validationError.ErrorID, errorMessage)
}
