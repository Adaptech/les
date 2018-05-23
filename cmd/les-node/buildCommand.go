package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"net/http"

	"github.com/Adaptech/les/pkg/eml"
	"github.com/Adaptech/les/pkg/eml/generate"
	"github.com/Adaptech/les/pkg/utils"
	"github.com/ghodss/yaml"
	"encoding/json"
	"strings"
)

const latestReleaseUrl = "https://api.github.com/repos/Adaptech/les-node-template/releases/latest"

func getLatestReleaseInfoFromGithub() (tagName string, browserDownloadUrl string) {
	tagName = ""
	browserDownloadUrl = ""

	metadata := map[string]interface{} {}
	if res, err := http.Get(latestReleaseUrl); err != nil {
		log.Printf("Error downloading latest release metadata: %s", err.Error())
		return
	} else if jsonMetadata, err := ioutil.ReadAll(res.Body); err != nil {
		log.Printf("Error reading latest release metadata: %s", err.Error())
		return
	} else if err := json.Unmarshal(jsonMetadata, &metadata); err != nil {
		log.Printf("Error parsing latest release metadata: %s", err.Error())
		return
	}

	var ok bool
	tagName, ok = metadata["tag_name"].(string)
	if !ok {
		log.Println("Error parsing latest release metadata: tag_name is not a string.")
		return
	}

	assets, ok := metadata["assets"].([]interface{})
	if !ok {
		log.Println("Error parsing latest release metadata: assets is not an array.")
		return
	}
	if len(assets) < 1 {
		log.Println("Error parsing latest release metadata: assets is empty.")
		return
	}
	asset, ok := assets[0].(map[string]interface{})
	if !ok {
		log.Println("Error parsing latest release metadata: asset[0] is not an object.")
	}
	browserDownloadUrl, ok = asset["browser_download_url"].(string)
	if !ok {
		log.Println("Error parsing latest release metadata: browser_download_url is not a string.")
		return
	}

	return
}

func ensureTemplateExists(templateDir string, lang string) {
	langTemplateDir := path.Join(templateDir, lang)

	fmt.Println("Checking for latest template version on github...")
	tagName, latestTemplateUrl := getLatestReleaseInfoFromGithub()

	var latestTemplateName string
	if tagName == "" || latestTemplateUrl == "" {
		fmt.Println("Looking for latest template version locally...")
		latestTemplateNameBytes, err := ioutil.ReadFile(path.Join(langTemplateDir, ".latest"))
		if err != nil {
			log.Fatalf("Error reading local latest metadata: %s", err.Error())
		}
		latestTemplateName = string(latestTemplateNameBytes)
	} else {
		latestTemplateName = strings.Replace(tagName, "release-", "les-node-template-", 1)
	}
	fmt.Println("Latest template version:", latestTemplateName)

	latestTemplateDir := path.Join(langTemplateDir, latestTemplateName)
	if _, err := os.Stat(latestTemplateDir); err == nil {
		return
	}

	fmt.Println(lang, "template", latestTemplateName, "missing locally, downloading from server...")
	if res, err := http.Get(latestTemplateUrl); err != nil {
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
