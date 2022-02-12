package pgen

import (
	"errors"
	"fmt"
	"strings"
	"text/template"
)

// Used to make certain string functions available in templates
var funcMap map[string]interface{} = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
}

func renderString(input string, vars map[string]interface{}) string {
	tmpl, err := template.New("").Funcs(funcMap).Parse(input)
	if err != nil {
		panic(err)
	}

	result := new(strings.Builder)
	err = tmpl.Execute(result, vars)
	if err != nil {
		panic(err)
	}

	return result.String()
}

func RenderTemplate(input *ProjectTemplate, userVars map[string]interface{}) (*RenderedTemplate, error) {

	if !validateUserDefinitions(input, userVars) {
		keys := make([]string, 0, len(userVars))
		for k := range userVars {
			keys = append(keys, k)
		}

		for _, v := range input.Variables {
			fmt.Printf("%s, ", v.Representation)
		}
		fmt.Println()

		for _, key := range keys {
			fmt.Printf("%s, ", key)
		}

		return nil, errors.New("User definitions do not match definitions expected by template")
	}

	var files []ProjectFile
	var dirs []string

	for _, f := range input.Files {
		files = append(files, ProjectFile{Path: renderString(f.Path, userVars), Content: renderString(f.Content, userVars)})
	}

	for _, d := range input.Directories {
		dirs = append(dirs, renderString(d, userVars))
	}

	tmpl := new(RenderedTemplate)
	tmpl.Files = files
	tmpl.Directories = dirs

	return tmpl, nil
}

func validateUserDefinitions(tmpl *ProjectTemplate, vars map[string]interface{}) bool {

	for _, v := range tmpl.Variables {
		if _, ok := vars[v.Identifier]; ok {
			continue
		} else {
			return false
		}
	}

	return true
}
