package main

import (
	"fmt"
	"path/filepath"

	"github.com/Adaptech/les/pkg/convert"
	"github.com/Adaptech/les/pkg/emd"
	"github.com/Adaptech/les/pkg/eml"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type convertCommand struct {
	conversionResult convert.EmdToEmlConversion
	file             *string
}

func configureConvertCommand(app *kingpin.Application) {
	c := &convertCommand{}
	convert := app.Command("convert", "Convert from Event Markdown (.emd) to Event Markup (.eml.yaml).").Action(c.convert)
	c.file = convert.Arg("file", ".eml.yaml file").String()
}

func (n *convertCommand) convert(c *kingpin.ParseContext) error {
	inputFile := useDefaultEmdFileIfInputFileNotSpecified(*n.file)
	if inputFile == "" {
		fmt.Println("No input file found. Try 'les convert --help'.")
		return nil
	}

	input, err := ReadLines(inputFile)
	if err != nil {
		return fmt.Errorf("convert command: %v", err)
	}
	markdown, err := emd.Parse(input)
	if err != nil {
		return fmt.Errorf("convert command: %v", err)
	}
	conversionResult, err := convert.EmdToEml(markdown)
	if err != nil {
		return fmt.Errorf("convert command: %v", err)
	}

	for _, markdownValidationError := range conversionResult.MarkdownValidationErrors {
		printError(markdownValidationError.ErrorID, markdownValidationError.Message)
	}
	for _, conversionError := range conversionResult.Eml.Errors {
		printError(conversionError.ErrorID, conversionError.Message)
	}

	if len(conversionResult.Eml.Errors) == 0 {
		n.conversionResult = conversionResult
		yaml, err := eml.ToYaml(n.conversionResult.Eml)
		if err != nil {
			return fmt.Errorf("convert yaml command: %v", err)
		}
		var extension = filepath.Ext(inputFile)
		var name = inputFile[0 : len(inputFile)-len(extension)]
		var outputFile = name + ".eml.yaml"

		err = WriteToFile(outputFile, yaml)
		if err != nil {
			return fmt.Errorf("convert command: %v", err)
		}
		fmt.Println("Input:\t\t" + inputFile)
		fmt.Println("Output:\t\t" + outputFile)
	}
	return nil
}
