package eml

import (
	"log"

	yaml "gopkg.in/yaml.v2"
)

// LoadYAML loads Event Markup Language YAML into a solution for further processing
func (c *Solution) LoadYAML(eml []byte) Solution {
	err := yaml.Unmarshal(eml, c)
	if err != nil {
		log.Printf("ERROR Invalid EML - could not parse %v", err)
	}

	return *c
}
