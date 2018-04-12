package generate

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func copyFile(src string, dst string) error {
	if fsrc, err := os.Open(src); err != nil {
		return err
	} else if fdst, err := os.Create(dst); err != nil {
		return err
	} else if _, err := io.Copy(fdst, fsrc); err != nil {
		return err
	}
	return nil
}

func copyDir(src string, dst string, recursive bool) error {
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, file := range files {
		var err error
		if file.IsDir() && recursive {
			err = copyDir(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()), true)
		} else {
			err = copyFile(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
		}
		if err != nil {
			return err
		}
	}
	return err
}

func copyInfrastructureTemplate(infrastructureTemplateDirectory string, renderingDirectory string) {
	ensureDirectoryExists(renderingDirectory)

	if err := copyDir(infrastructureTemplateDirectory, renderingDirectory, true); err != nil {
		log.Fatalf("copyInfrastructureTemplate failed: %v", err)
	}
}

func writeRenderedEvent(eventsDirectory string, eventName string, code string, ext string) {
	ensureDirectoryExists(eventsDirectory)

	err := ioutil.WriteFile(filepath.Join(eventsDirectory, eventName+ext), []byte(code), 0644)
	if err != nil {
		log.Fatalf("writeRenderedEvent: %v", err)
	}
}

func writeRenderedReadmodel(readmodelDirectory string, readmodelName string, code string, ext string) {
	ensureDirectoryExists(readmodelDirectory)

	err := ioutil.WriteFile(filepath.Join(readmodelDirectory, readmodelName+ext), []byte(code), 0644)
	if err != nil {
		log.Fatalf("writeRenderedReadmodel: %v", err)
	}
}

func writeRenderedCommand(commandsDirectory, commandName string, code string, ext string) {
	ensureDirectoryExists(commandsDirectory)

	err := ioutil.WriteFile(filepath.Join(commandsDirectory, commandName+ext), []byte(code), 0644)
	if err != nil {
		log.Fatalf("writeRenderedCommand: %v", err)
	}

}
func writeRenderedAggregate(domainDirectory string, aggregateName string, code string, ext string) {
	ensureDirectoryExists(domainDirectory)

	err := ioutil.WriteFile(filepath.Join(domainDirectory, aggregateName+ext), []byte(code), 0644)
	if err != nil {
		log.Fatalf("writeRenderedAggregate: %v", err)
	}
}

func writeRenderedController(controllerDirectory string, controllerName string, code string, ext string) {
	ensureDirectoryExists(controllerDirectory)

	err := ioutil.WriteFile(filepath.Join(controllerDirectory, controllerName+ext), []byte(code), 0644)
	if err != nil {
		log.Fatalf("writeRenderedAggregate: %v", err)
	}
}

func ensureDirectoryExists(directory string) {
	if err := os.MkdirAll(directory, 0755); err != nil {
		log.Fatalf("ensureDirectoryExists failed: %v", err)
	}

}
