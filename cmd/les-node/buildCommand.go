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

const latestTemplateName = "les-node-template-20180503-19cddd0"
const nodejsTemplateDownloadURL = "https://github.com/Adaptech/les/raw/master/releases/les-node-template/" + latestTemplateName + ".tar.gz"

func ensureTemplateExists(templateDir string, lang string) {
	langTemplateDir := path.Join(templateDir, lang)
	latestTemplateDir := path.Join(langTemplateDir, latestTemplateName)
	if _, err := os.Stat(latestTemplateDir); err == nil {
		return
	}

	var templateURL = nodejsTemplateDownloadURL
	fmt.Println(lang, "template missing locally, downloading from server...")
	if res, err := http.Get(templateURL); err != nil {
		log.Fatalf("Error downloading template: %s", err.Error())
	} else if err := utils.Untar(langTemplateDir, res.Body); err != nil {
		log.Fatalf("Error extracting archive: %s", err.Error())
	}
	if err := ioutil.WriteFile(path.Join(langTemplateDir, ".latest"), []byte(latestTemplateName), 0644); err != nil {
		log.Fatalf("Error creating .latest file: %s", err.Error())
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
