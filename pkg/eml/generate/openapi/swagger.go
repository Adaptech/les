package openapi

import (
	"bytes"
	"log"

	"github.com/Adaptech/les/pkg/eml/generate/nodejs"

	"github.com/alecthomas/template"
	"github.com/Adaptech/les/pkg/eml"
)

// SwaggerYML ...
func SwaggerYML(solution eml.Solution) string {
	const tpl = `openapi: "3.0.1"

info:
  title: "{{.Solution.Name}}"
  description: "{{.Solution.Name}} API"
  version: "0.1"

servers:
  - url: http://localhost:3001/api/v1
    description: localhost

paths:
    {{range $cnt, $context := $.Solution.Contexts}}{{range $cnt, $stream := $context.Streams}}{{range $cnt, $command := $stream.Commands}}/{{$stream.Name}}/{{$command.Command.Name | ToNodeJsClassName }}:
      post:
        tags: 
          - {{$stream.Name}}
        summary: {{$command.Command.Name}}
        description: {{$command.Command.Name}} command.
        requestBody:
          required: true
          content:
            application/json:
              schema:
                type: object
                properties:
                  {{range $cnt, $parameter := $command.Command.Parameters}}{{$parameter.Name}}:
                    type: {{$parameter.Type}}
                    example: "{{ (call $.Data.Parameter $parameter.Name $stream.Name)}}"
                  {{end}}
        responses:
          '202':
            description: "Accepted."
          '400':
            description: "Business rule validation failed. Could not execute command."
          default:
            description: "Unexpected error."
    {{end}}{{end}}{{end}}
    {{range $cnt, $context := $.Solution.Contexts}}{{range $cnt, $readmodel := $context.Readmodels}}
    /r/{{$readmodel.Readmodel.Name | ToNodeJsClassName }}:
      get:
        tags: 
          - Queries
        description: {{$readmodel.Readmodel.Name}}
        responses:
          '200':
            description: OK
          default:
            description: "Unexpected error."
    {{end}}{{end}}`

	funcMap := template.FuncMap{
		"ToNodeJsClassName": nodejs.ToNodeJsClassName,
	}
	t := template.Must(template.New("eml").Funcs(funcMap).Parse(tpl))
	buf := bytes.NewBufferString("")

	example := Example{
		Parameter: parameterGenerator,
	}
	data := SwaggerData{
		Solution: solution,
		Data:     example,
	}

	err := t.Execute(buf, data)
	if err != nil {
		log.Fatal("error executing SwaggerYML template:", err)
	}
	return buf.String()

}

// SwaggerData ...
type SwaggerData struct {
	Solution eml.Solution
	Data     Example
}
