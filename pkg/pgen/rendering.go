package pgen

import (
	"errors"
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

func RenderTemplate(input ProjectTemplate, userVars map[string]interface{}) (RenderedTemplate, error) {

	for _, v := range input.Variables {
		if _, ok := userVars[v]; ok {
			continue
		} else {
			return RenderedTemplate{}, errors.New("missing variable definition")
		}
	}

	var files []ProjectFile
	var dirs []string

	for _, f := range input.Files {
		files = append(files, ProjectFile{Path: renderString(f.Path, userVars), Content: renderString(f.Content, userVars)})
	}

	for _, d := range input.Directories {
		dirs = append(dirs, renderString(d, userVars))
	}

	tmpl := RenderedTemplate{
		Files:       files,
		Directories: dirs,
	}

	return tmpl, nil
}
