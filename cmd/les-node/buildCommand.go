package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"net/http"

	"github.com/ghodss/yaml"
	"github.com/Adaptech/les/pkg/eml"
	"github.com/Adaptech/les/pkg/eml/generate"
	"github.com/Adaptech/les/pkg/utils"
)

const nodejsTemplateDownloadURL = "https://esp-templates.azurewebsites.net/esp-nodejs-template-latest.tar.gz"

func ensureTemplateExists(templateDir string, lang string) {
	langTemplateDir := path.Join(templateDir, lang)
	if _, err := os.Stat(langTemplateDir); err == nil {
		return
	}

	var templateURL = nodejsTemplateDownloadURL
	fmt.Println(lang, "template missing locally, downloading from server...")
	if res, err := http.Get(templateURL); err != nil {
		log.Fatalf("Error downloading template: %s", err.Error())
	} else if err := utils.Untar(templateDir, res.Body); err != nil {
		log.Fatalf("Error extracting archive: %s", err.Error())
	}
}

func writeSwaggerJSONFile(markup eml.Solution, buildDir string) {
	swaggerYAML := generate.OpenAPISpec(markup)
	err := ioutil.WriteFile(filepath.Join(buildDir, "swagger.yaml"), []byte(swaggerYAML), 0644)
	if err != nil {
		log.Fatalf("writeSwaggerJSONFile WriteFile: Error writing swagger.json - %v", err)
	}

	swaggerJSON, err := yaml.YAMLToJSON([]byte(swaggerYAML))
	if err != nil {
		log.Fatalf("writeSwaggerJSONFile YAMLToJSON: %v", err)
	}
	err = ioutil.WriteFile(filepath.Join(buildDir, "swagger.json"), []byte(swaggerJSON), 0644)
	if err != nil {
		log.Fatalf("writeSwaggerJSONFile WriteFile: Error writing swagger.json - %v", err)
	}
}

func buildAPIFrom(inputFile string) error {
	fileContent, err := ReadFile(inputFile)
	if err != nil {
		return err
	}
	markup := eml.Solution{}
	markup.LoadYAML(fileContent)
	espdirectory, err := getDirectories()
	if err != nil {
		log.Fatal(err)
	}
	buildDir := espdirectory.build
	_ = os.MkdirAll(buildDir, os.ModePerm)
	templateDir := espdirectory.template

	_ = os.MkdirAll(buildDir, os.ModePerm)
	ensureTemplateExists(templateDir, "nodejs")
	generate.NodeAPI(markup, buildDir, templateDir)
	writeSwaggerJSONFile(markup, buildDir)
	fmt.Println("Source Code:\t./" + buildDir)
	fmt.Println("URI:\t\thttp://localhost:3001/api/v1")
	fmt.Println("Eventstore DB:\thttp://localhost:2113 (username 'admin', password 'changeit')")
	fmt.Println("API Docs:\thttp://localhost:3001/api-docs")
	fmt.Printf("Start API:\tcd api && npm install && docker-compose up -d\n")
	return nil
}
