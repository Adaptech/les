package nodejs

import (
	"bytes"
	"html/template"
	"log"
	"sort"

	"github.com/Adaptech/les/pkg/eml"
)

// EventToJs renders nodeJs for an event
func EventToJs(event eml.Event) string {
	const eventTemplate = `export default class {{ .Event.Name | ToNodeJsClassName }} {
  constructor({{range $cnt, $property := $.Event.Properties}}{{if gt $cnt 0}}, {{end}}{{$property.Name}}{{end}}) {
		{{range $cnt, $property := $.Event.Properties}}this.{{$property.Name}} = {{$property.Name}};
		{{end}}
  }
}
`
	sort.Slice(event.Event.Properties, func(i, j int) bool { return event.Event.Properties[i].Name < event.Event.Properties[j].Name })

	funcMap := template.FuncMap{
		"ToNodeJsClassName": ToNodeJsClassName,
	}

	t := template.Must(template.New("event").Funcs(funcMap).Parse(eventTemplate))
	buf := bytes.NewBufferString("")
	err := t.Execute(buf, event)
	if err != nil {
		log.Fatal("error executing EventToJS template:", err)
	}
	return buf.String()
}
