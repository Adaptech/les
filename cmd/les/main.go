package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

const defaultEmdFile = "Eventmarkdown.emd"
const defaultEmlFile = "Emlfile.yaml"
const generatedEmlFile = ".generated.eml.yaml"

func main() {
	app := kingpin.New("les", "Let's Event Source: Validate & convert Event Markup Language and Event Markup.")
	configureConvertCommand(app)
	configureValidateCommand(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func printError(id string, message string) {
	fmt.Printf("%s: %s\n", id, message)
}
