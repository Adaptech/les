package main

import (
	"fmt"
	"log"

	"github.com/Adaptech/les/pkg/convert"
	"github.com/Adaptech/les/pkg/emd"
	"github.com/Adaptech/les/pkg/eml"
)

func convertFileToEml(inputFile string, outputFile string) (bool, error) {
	input, err := ReadLines(inputFile)
	if err != nil {
		return false, fmt.Errorf("convertToEml: %v", err)
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		return false, fmt.Errorf("convertToEml: %v", err)
	}
	conversionResult, err := convert.EmdToEml(markdown)
	if err != nil {
		return false, fmt.Errorf("convertToEml: %v", err)
	}

	isValidEml := len(conversionResult.MarkdownValidationErrors) == 0
	if !isValidEml {
		fmt.Println("EMD Errors:")
		for _, emdErr := range conversionResult.MarkdownValidationErrors {
			printError(emdErr.ErrorID, emdErr.Message)
		}
	}

	markup := conversionResult.Eml
	yaml, err := eml.ToYaml(markup)
	if err != nil {
		log.Panicf("convertFileToEml: %v", err)
	}
	err = WriteToFile(outputFile, yaml)
	if err != nil {
		log.Panicf("convertFileToEml: %v", err)
	}
	return isValidEml, nil
}
