package generate

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Adaptech/les/pkg/eml"
	"github.com/Adaptech/les/pkg/eml/generate/nodejs"
)

const nodeJsExt = ".js"

// NodeAPI Event Markup Language to a nodejs API
func NodeAPI(system eml.Solution, renderingDirectory string, infrastructureTemplateDirectory string) {
	var nodeJsRenderingDirectory = renderingDirectory
	var nodeJsTemplateDirectory = filepath.Join(infrastructureTemplateDirectory, "nodejs")
	var eventsDirectory = filepath.Join(nodeJsRenderingDirectory, "src", "events")
	var commandsDirectory = filepath.Join(nodeJsRenderingDirectory, "src", "commands")
	var domainDirectory = filepath.Join(nodeJsRenderingDirectory, "src", "domain")
	var controllerDirectory = filepath.Join(nodeJsRenderingDirectory, "src", "controllers")
	var readmodelDirectory = filepath.Join(nodeJsRenderingDirectory, "src", "readModels")

	deleteAllExceptNodeModules(nodeJsRenderingDirectory)
	copyInfrastructureTemplate(nodeJsTemplateDirectory, nodeJsRenderingDirectory)

	for _, boundedContext := range system.Contexts {
		for _, stream := range boundedContext.Streams {
			for _, event := range stream.Events {
				renderedJavascript := nodejs.EventToJs(event)
				writeRenderedEvent(eventsDirectory, nodejs.ToNodeJsClassName(event.Event.Name), renderedJavascript, nodeJsExt)
			}
			for _, command := range stream.Commands {
				renderedJavascript := nodejs.CommandToJs(command)
				writeRenderedCommand(commandsDirectory, nodejs.ToNodeJsClassName(command.Command.Name), renderedJavascript, nodeJsExt)
			}
			renderedJavaScript := nodejs.DomainJs(stream, stream.Events)
			writeRenderedAggregate(domainDirectory, stream.Name, renderedJavaScript, nodeJsExt)

			renderedJavaScript = nodejs.ControllerJs(stream, readmodelLookupFor(boundedContext))
			writeRenderedController(controllerDirectory, stream.Name, renderedJavaScript, nodeJsExt)

		}
		eventLookup := make(map[string]eml.Event)
		for _, stream := range boundedContext.Streams {
			for _, event := range stream.Events {
				eventLookup[event.Event.Name] = event
			}
		}
		for _, readmodel := range boundedContext.Readmodels {
			readModelJs := nodejs.ReadmodelsToJs(readmodel, eventLookup)
			writeRenderedReadmodel(readmodelDirectory, nodejs.ToNodeJsClassName(readmodel.Readmodel.Name), readModelJs, nodeJsExt)
		}
	}
}

func deleteAllExceptNodeModules(nodeJsRenderingDirectory string) {
	files, err := filepath.Glob(filepath.Join(nodeJsRenderingDirectory, "*.*"))
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
	if err := os.RemoveAll(filepath.Join(nodeJsRenderingDirectory, "Dockerfile")); err != nil {
		log.Fatalf("NodeAPI deletePreviouslyRendered failed: %v", err)
	}
	if err := os.RemoveAll(filepath.Join(nodeJsRenderingDirectory, "config")); err != nil {
		log.Fatalf("NodeAPI deletePreviouslyRendered failed: %v", err)
	}
	if err := os.RemoveAll(filepath.Join(nodeJsRenderingDirectory, "src")); err != nil {
		log.Fatalf("NodeAPI deletePreviouslyRendered failed: %v", err)
	}
	if err := os.RemoveAll(filepath.Join(nodeJsRenderingDirectory, "test")); err != nil {
		log.Fatalf("NodeAPI deletePreviouslyRendered failed: %v", err)
	}
	if err := os.RemoveAll(filepath.Join(nodeJsRenderingDirectory, "web")); err != nil {
		log.Fatalf("NodeAPI deletePreviouslyRendered failed: %v", err)
	}
}

func readmodelLookupFor(boundedContext eml.BoundedContext) map[string]eml.Readmodel {
	readmodelLookup := make(map[string]eml.Readmodel)
	for _, readmodel := range boundedContext.Readmodels {
		readmodelLookup[readmodel.Readmodel.Name] = readmodel
	}
	return readmodelLookup
}
