package pgen

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

func LoadFromJson(data string) (ProjectTemplate, error) {

	tmpl := ProjectTemplate{}
	err := json.Unmarshal([]byte(data), &tmpl)
	return tmpl, err
}

func LoadFromYaml(data string) (ProjectTemplate, error) {
	tmpl := ProjectTemplate{}
	err := yaml.Unmarshal([]byte(data), &tmpl)
	return tmpl, err
}
