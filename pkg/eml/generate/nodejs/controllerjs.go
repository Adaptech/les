package nodejs

import (
	"bytes"
	"html/template"
	"log"
	"strings"

	"github.com/Adaptech/les/pkg/eml"
)

// ControllerJs renders a REST API controller.
func ControllerJs(stream eml.Stream, readmodelLookup map[string]eml.Readmodel) string {
	const controllerTemplate = `import {{ .Stream.Name | ToNodeJsClassName }} from '../domain/{{ .Stream.Name | ToNodeJsClassName }}';
{{range $cnt, $command := $.Stream.Commands}}import {{$command.Command.Name | ToNodeJsClassName }} from '../commands/{{$command.Command.Name | ToNodeJsClassName}}';
{{end}}

export default class {{ .Stream.Name }}Controller {
  constructor(app, readRepository, commandHandler, logger) {
	{{range $cnt, $command := $.Stream.Commands}}
    async function {{$command.Command.Name | ToNodeJsClassName | ToLower}}(req, res) {
			let { {{range $cnt, $parameter := $command.Command.Parameters}}{{if gt $cnt 0}}, {{end}}{{$parameter.Name}}{{end}} } = req.body;
			let foundItem = null;
			{{range $cnt, $parameter := $command.Command.Parameters}}{{if eq ($parameter.RuleExists "MustExistIn") true }}
			foundItem = await readRepository.findOne('{{ $parameter.MustExistInReadmodel }}', { {{ $parameter.MustExistInReadmodel | GetReadmodelKey }}: { eq: req.body.{{$parameter.Name}} } }, true);
			{{$parameter.Name}} = foundItem && foundItem.{{ $parameter.MustExistInReadmodel | GetReadmodelKey }};{{end}}			{{end}}
      const command = new {{$command.Command.Name | ToNodeJsClassName }}({{range $cnt, $parameter := $command.Command.Parameters}}{{if gt $cnt 0}}, {{end}}{{$parameter.Name}}{{end}});
      commandHandler(command.{{ $.Stream.Name | ToNodeJsClassName | ToLower }}Id, new {{$.Stream.Name | ToNodeJsClassName}}(), command)
          .then(() => {
            res.status(202).json(command);
          })
          .catch(err => {
            if(err.name == "ValidationFailed") {
              res.status(400).json({message: err.message});
            } else {
              logger.error(err.stack);
              res.status(500).json({message: err.message});
            }
          });
		}
    app.post('/api/v1/{{ $.Stream.Name }}/{{$command.Command.Name | ToNodeJsClassName }}', {{$command.Command.Name | ToNodeJsClassName | ToLower}});
		{{end}}
	}
}
`
	ReadmodelKeyLookup := func(modelName string) string {
		return readmodelLookup[modelName].Readmodel.Key
	}

	funcMap := template.FuncMap{
		"ToLower":           strings.ToLower,
		"ToNodeJsClassName": ToNodeJsClassName,
		"GetReadmodelKey":   ReadmodelKeyLookup,
	}

	t := template.Must(template.New("controller").Funcs(funcMap).Parse(controllerTemplate))

	type templateData struct {
		Stream    eml.Stream
		Readmodel map[string]eml.Readmodel
	}

	data := templateData{
		Stream: stream,
	}

	buf := bytes.NewBufferString("")
	err := t.Execute(buf, data)
	if err != nil {
		log.Fatal("error executing ControllerJS template:", err)
	}
	return buf.String()

}
