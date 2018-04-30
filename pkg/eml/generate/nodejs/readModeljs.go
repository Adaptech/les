package nodejs

import (
	"bytes"
	"html/template"
	"log"

	"github.com/Adaptech/les/pkg/eml"
)

// ReadmodelsToJs renders nodeJs for a command
func ReadmodelsToJs(readmodel eml.Readmodel, eventLookup map[string]eml.Event) string {
	const readmodelTemplate = `export const config = {
  key: '{{$.Readmodel.Readmodel.Key}}',
  schema: {
    {{$.Readmodel.Readmodel.Key}}: {type: 'string', nullable: false},
    {{ range $cnt, $property := $.Properties}}{{ $property.Name }}: {type: '{{$property.Type}}'},
    {{end}}
  }
};

export async function handler({{ $.Readmodel.Readmodel.Name | ToNodeJsClassName }}Repo, eventData, lookups) {
  const { typeId, event } = eventData;
  switch (typeId) {
    {{ range $cnt, $event := $.Readmodel.Readmodel.SubscribesTo  }}case '{{$event}}': {
      await createOrUpdate({{ $.Readmodel.Readmodel.Name | ToNodeJsClassName }}Repo, { {{$.Readmodel.Readmodel.Key}}: event.{{$.Readmodel.Readmodel.Key}} }, {
        {{ range $cnt, $property := $.Properties}}{{if eq $property.Event $event}}{{$property.Name}}: event.{{$property.Name}},
        {{end}}{{end}}
      });
      break;
    }
    {{end}}
  }
  return {{ $.Readmodel.Readmodel.Name | ToNodeJsClassName }}Repo;
}

async function createOrUpdate(repo, idQuery, data) {
  const exists = await repo.exists(idQuery);
  if (!exists) {
    const payload = {...idQuery, ...data};
    for (const k of Object.keys(config.schema)) {
      if (!payload[k]) payload[k] = "";
    }
    repo.create(payload);
  } else {
    const payload = {...data};
    repo.updateOne(idQuery, payload);
  }
}
`

	type TemplateData struct {
		Readmodel  eml.Readmodel
		Properties map[string]struct {
			Name  string
			Type  string
			Event string
		}
		Key string
	}

	hasPreconditionEvent := make(map[string]bool)
	for _, eventID := range readmodel.Readmodel.SubscribesTo {
		hasPreconditionEvent[eventID] = true
	}

	properties := make(map[string]struct {
		Name  string
		Type  string
		Event string
	})

	// Find what properties belong to what events so that they can be populated from events the read model subscribes to:
	for _, event := range eventLookup {
		if _, ok := hasPreconditionEvent[event.Event.Name]; ok {
			for _, property := range event.Event.Properties {
				if property.Name != readmodel.Readmodel.Key {
					properties[property.Name] = struct {
						Name  string
						Type  string
						Event string
					}{
						Name:  property.Name,
						Type:  property.Type,
						Event: event.Event.Name,
					}
				}
			}
		}
	}

	data := TemplateData{
		Readmodel:  readmodel,
		Properties: properties,
	}

	funcMap := template.FuncMap{
		"ToNodeJsClassName": ToNodeJsClassName,
	}

	t := template.Must(template.New("readmodel").Funcs(funcMap).Parse(readmodelTemplate))
	buf := bytes.NewBufferString("")
	err := t.Execute(buf, data)
	if err != nil {
		log.Fatal("error executing ReadModelToJS template:", err)
	}
	return buf.String()
}
