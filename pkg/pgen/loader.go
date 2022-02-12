package pgen

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

func LoadFromJson(data []byte) (*ProjectTemplate, error) {

	tmpl := new(ProjectTemplate)
	err := json.Unmarshal(data, &tmpl)
	return tmpl, err
}

func LoadFromYaml(data []byte) (*ProjectTemplate, error) {
	tmpl := new(ProjectTemplate)
	err := yaml.Unmarshal(data, &tmpl)
	return tmpl, err
}
