package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/radovskyb/watcher"
)

func whenFileChangesThenValidate(emdFile string, folder string) {
	w := watcher.New()
	w.FilterOps(watcher.Write)
	w.IgnoreHiddenFiles(true)

	go func() {
		for {
			select {
			case event := <-w.Event:
				processEvent(event)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add(folder); err != nil {
		log.Fatalln(err)
	}

	// Initial validation on startup:
	for _, f := range w.WatchedFiles() {
		if strings.HasSuffix(f.Name(), emdFile) || strings.HasSuffix(f.Name(), defaultEmlFile) {
			fmt.Println("Initial validation: " + f.Name())
			validateFile(f.Name())
		}
	}

	go func() {
		w.Wait()
		w.TriggerEvent(watcher.Write, nil)
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

func processEvent(event watcher.Event) {
	if !event.IsDir() && len(event.Path) >= 4 {
		validateFile(event.Path)
	}
}

func validateFile(fileName string) {
	if len(fileName) >= 4 {
		isValidatingEmdFile := strings.HasSuffix(fileName, ".emd")
		isValidatingEmlFile := strings.HasSuffix(fileName, ".eml.yaml")
		if isValidatingEmdFile {
			isValidEmd, err := convertFileToEml(fileName, generatedEmlFile)
			if err != nil {
				log.Panicln(err)
			}
			isValidEml, err := checkIfFileContainsValidEml(generatedEmlFile)
			if err != nil {
				log.Panicln(err)
			}
			if isValidEml && isValidEmd {
				fmt.Println("OK")
			}
		}
		if isValidatingEmlFile {
			isValidEmd, err := checkIfFileContainsValidEml(fileName)
			if err != nil {
				log.Panicln(err)
			}
			if isValidEmd {
				fmt.Println("OK")
			}
		}
	}
}
