package nodejs

import (
	"bytes"
	"html/template"
	"log"
	"strings"

	"github.com/Adaptech/les/pkg/eml"
)

// Field for domain templates
type Field struct {
	Name     string
	Type     string
	IsHashed bool
}

// DomainTemplateParams ...
type DomainTemplateParams struct {
	Stream              eml.Stream
	EventLookup         map[string][]eml.Property
	Fields              map[string]Field
	HasHashedProperties bool
}

// DomainJs renders an aggregate for an event stream.
func DomainJs(stream eml.Stream, eventList []eml.Event) string {
	const aggregateTemplate = `{{range $cnt, $command := $.Stream.Commands}}import {{$command.Command.Name | ToNodeJsClassName}} from '../commands/{{$command.Command.Name | ToNodeJsClassName}}';
{{end}}{{range $cnt, $event := $.Stream.Events}}import {{$event.Event.Name  | ToNodeJsClassName }} from '../events/{{$event.Event.Name | ToNodeJsClassName}}';
{{end}}import errors from '../domain/errors';
{{if eq .HasHashedProperties true}}import bcrypt from 'bcrypt';	
{{end}}
export default class {{ .Stream.Name }} {
  constructor() {
    this._id = null;
    this._eventCount = new Map();
  }

  getEventCount(eventType) {
    const previousOccurrencesCount = this._eventCount.get(eventType);
    if (previousOccurrencesCount) {
      return previousOccurrencesCount;
    }
    return 0;
  }

  incrementEventOccurenceCount(eventType) {
    const previousOccurrencesCount = this._eventCount.get(eventType);
    if (previousOccurrencesCount) {
      this._eventCount.set(eventType, previousOccurrencesCount + 1);
    } else {
      this._eventCount.set(eventType, 1);
    }
  }
	
  hydrate(evt) {
	{{range $cnt, $event := $.Stream.Events}}if (evt instanceof {{$event.Event.Name | ToNodeJsClassName}}) {
      this.incrementEventOccurenceCount("{{$event.Event.Name | ToNodeJsClassName}}");
      this._on{{$event.Event.Name | ToNodeJsClassName}}(evt);
    }
	{{end}}
  }
  
  {{range $cnt, $event := $.Stream.Events}}
  _on{{$event.Event.Name | ToNodeJsClassName}}(evt) {
    {{range $cnt, $property := $event.Event.Properties}}this._{{$property.Name}} = evt.{{$property.Name}};
    {{end}}
  }
  {{end}}

  execute(command) {
    {{range $cnt, $command := $.Stream.Commands}}if (command instanceof {{$command.Command.Name | ToNodeJsClassName}}) {
      return this._{{$command.Command.Name | ToNodeJsClassName}}(command);
    }
    {{end}}
    throw new Error('Unknown command.');
  }

  {{range $cnt, $command := $.Stream.Commands}}
  {{if eq $.HasHashedProperties true}}async {{end}}_{{$command.Command.Name | ToNodeJsClassName}}(command) {
    const validationErrors = [];
    {{range $cnt, $parameter := $command.Command.Parameters}}{{if eq (.RuleExists "MustExistIn") true}}if (!command.{{$parameter.Name}}) {
      validationErrors.push({ field: "{{$parameter.Name}}", msg: "{{$parameter.Name}} does not exist." });
    }{{end}}{{end}}
    {{range $cnt, $parameter := $command.Command.Parameters}}{{if eq (.RuleExists "IsRequired") true}}if (!command.{{$parameter.Name}}) {
      validationErrors.push({ field: "{{$parameter.Name}}", msg: "{{$parameter.Name}} is a required field." });
    }{{end}}{{end}}
    {{range $cnt, $precondition := $command.Command.Preconditions}}
    {{if eq "MustHaveHappened" (GetToken $precondition 1) }}if (this.getEventCount("{{GetToken $precondition 0}}") == 0) {
      // {{$precondition}}
      validationErrors.push({ field: "", msg: "Cannot {{$command.Command.Name}}. {{GetToken $precondition 0}} must have occurred." });
    }{{end}}
    {{if eq "MustNotHaveHappened" (GetToken $precondition 1) }}if (this.getEventCount("{{GetToken $precondition 0}}") != 0) {
      // {{$precondition}}
      validationErrors.push({ field: "", msg: "Cannot {{$command.Command.Name}}. {{GetToken $precondition 0}} must not have occurred." });
    }{{end}}
    {{end}}
    if (validationErrors.length > 0) {
      throw new errors.ValidationFailed(validationErrors);
    }
    {{range $cnt, $postcondition := $command.Command.Postconditions}}{{range $cnt, $parameter := index $.EventLookup $postcondition }}{{if eq $parameter.IsHashed true}}command.{{$parameter.Name}} = await new Promise((resolve, reject) => bcrypt.hash(command.{{$parameter.Name}}, 10, function(err, hash) {
      if (err) reject(err);
      else resolve(hash);
    }));{{end}}{{end}}{{end}}
    const result = [];{{range $cnt, $postcondition := $command.Command.Postconditions}}
    const event{{$cnt}} = new {{ $postcondition | ToNodeJsClassName }}();
    {{range $param_cnt, $parameter := $command.Command.Parameters}}event{{$cnt}}.{{$parameter.Name}} = command.{{$parameter.Name}} || "";
    {{end}}
    result.push(event{{$cnt}});
    {{end}}
    return result;
  }
  {{end}}
}
`
	// {{range $cnt, $parameter := index $.EventLookup $postcondition }}{{if gt $cnt 0}}, {{end}}command.{{$parameter.Name}}{{end}}
	// {{range $cnt, $parameter := $command.Command.Parameters}}event.command.{{$parameter.Name}} = command.{{$parameter.Name}};{{end}}
	eventLookup := make(map[string][]eml.Property)
	for _, event := range eventList {
		eventLookup[event.Event.Name] = event.Event.Properties
	}

	fields := map[string]Field{}
	hasHashedEventProperty := false
	for _, event := range eventList {
		for _, prop := range event.Event.Properties {
			fields[prop.Name] = Field{prop.Name, prop.Type, prop.IsHashed}
			if prop.IsHashed == true {
				hasHashedEventProperty = true
			}
		}
	}

	domainTemplateData := DomainTemplateParams{
		Stream:              stream,
		EventLookup:         eventLookup,
		Fields:              fields,
		HasHashedProperties: hasHashedEventProperty,
	}

	funcMap := template.FuncMap{
		"ToNodeJsClassName": ToNodeJsClassName,
		"GetToken":          getToken,
		"Contains":          strings.Contains,
	}

	t := template.Must(template.New("aggregate").Funcs(funcMap).Parse(aggregateTemplate))
	buf := bytes.NewBufferString("")
	err := t.Execute(buf, domainTemplateData)
	if err != nil {
		log.Fatal("error executing DomainJs template:", err)
	}
	return buf.String()
}
