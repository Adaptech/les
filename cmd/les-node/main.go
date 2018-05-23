package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

const defaultEmlFile = "Eventstorming.eml.yaml"

var (
	buildAPI  = kingpin.Flag("build", "Build a NodeJS API from EML.").Short('b').Bool()
	inputFile = kingpin.Arg("file", ".eml.yaml file. Default: "+defaultEmlFile).String()
)

func main() {
	kingpin.Version("0.10.6-alpha")
	kingpin.Parse()
	if *buildAPI {
		inputFile := useDefaulEmlFileIfInputFileNotSpecified(*inputFile)
		if inputFile == "" {
			fmt.Println("No input file found. Try 'les-node --help'.")
			os.Exit(-1)
		}
		fmt.Println("API Spec:\t" + inputFile)
		isValidEml, err := checkIfFileContainsValidEml(inputFile)
		if err != nil {
			log.Fatal("build:", err)
		}
		if isValidEml {
			buildAPIFrom(inputFile)
		}
	} else {
		kingpin.Usage()
	}
	os.Exit(0)
}

func printError(id string, message string) {
	fmt.Printf("%s: %s\n", id, message)
}
