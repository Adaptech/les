package main

import (
	"fmt"

	"github.com/Adaptech/les/pkg/emd"
)

func emdToGraphVizDigraph(inputFile string) error {
	if inputFile == "" {
		fmt.Println("No input file found. Try 'les-viz --help'.")
		return nil
	}

	eventstorming, err := ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("les-viz: %v", err)
	}
	graphVizDigraph := emd.ToGraphViz(string(eventstorming))
	fmt.Println(graphVizDigraph)
	return nil
}
