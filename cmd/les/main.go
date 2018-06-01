package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

const defaultEmdFile = "Eventstorming.emd"
const defaultEmlFile = "Eventstorming.eml.yaml"
const generatedEmlFile = ".generated.eventsourcing.eml.yaml"

func main() {
	app := kingpin.New("les", "Event Markdown (EMD) and Event Markup Language (EML) validation and conversion.")
	app.Version("0.10.6-alpha")
	configureConvertCommand(app)
	configureValidateCommand(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func printError(id string, message string) {
	fmt.Printf("%s: %s\n", id, message)
}
