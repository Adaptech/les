package nodejs

import (
	"bytes"
	"html/template"
	"log"
	"sort"

	"github.com/Adaptech/les/pkg/eml"
)

// CommandToJs renders nodeJs for a command
func CommandToJs(command eml.Command) string {
	const commandTemplate = `export default class {{ .Command.Name | ToNodeJsClassName }} {
	constructor({{range $cnt, $parameter := $.Command.Parameters}}{{if gt $cnt 0}}, {{end}}{{$parameter.Name}}{{end}}) {
		{{range $cnt, $parameter := $.Command.Parameters}}this.{{$parameter.Name}} = {{$parameter.Name}};
		{{end}}
  }
}
`
	sort.Slice(command.Command.Parameters, func(i, j int) bool { return command.Command.Parameters[i].Name < command.Command.Parameters[j].Name })

	funcMap := template.FuncMap{
		"ToNodeJsClassName": ToNodeJsClassName,
	}

	t := template.Must(template.New("command").Funcs(funcMap).Parse(commandTemplate))
	buf := bytes.NewBufferString("")
	err := t.Execute(buf, command)
	if err != nil {
		log.Fatal("error executing CommandToJs template:", err)
	}
	return buf.String()
}
