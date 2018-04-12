package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type validateCommand struct {
	tail      *bool
	inputType *string
	file      *string
}

func configureValidateCommand(app *kingpin.Application) {
	c := &validateCommand{}
	cmd := app.Command("validate", "Verify if event markdown in ./"+defaultEmdFile+" or event markup in ./"+defaultEmlFile+" is valid for building APIs.")
	c.tail = cmd.Flag("follow", "Re-validate whenever ./"+defaultEmdFile+" or  ./"+defaultEmlFile+" are changed.").Short('f').Bool()
	c.file = cmd.Arg("file", ".emd or .eml.yaml file. Default: "+defaultEmdFile+" or "+defaultEmlFile+".").String()
	cmd.Action(c.validate)
}

func (n *validateCommand) validate(c *kingpin.ParseContext) error {
	inputFile := useDefaultEmdOrEmlFileIfInputFileNotSpecified(*n.file)
	if inputFile == "" {
		fmt.Println("No input file found. Try 'les validate --help'.")
		return nil
	}
	tail := false
	if n.tail != nil {
		tail = *n.tail
	}
	if tail {
		whenFileChangesThenValidate(inputFile, ".")
		return nil
	}

	validateFile(inputFile)
	return nil
}
