package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

const defaultEmdFile = "Eventstorming.emd"

func main() {
	app := kingpin.New("les-viz", "Generates a http://www.graphviz.org/ digraph for an event storming.")
	app.Version("0.10.6-alpha")
	file := app.Arg("file", "Event Markdown (.emd) file").String()
	kingpin.MustParse(app.Parse(os.Args[1:]))
	emdToGraphVizDigraph(*file)
}

func printError(id string, message string) {
	fmt.Printf("%s: %s\n", id, message)
}
