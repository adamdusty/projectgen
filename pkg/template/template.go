package template

import (
	"errors"
	"strings"
	"text/template"
)

type ProjectFile struct {
	Path    string
	Content string
}

type ProjectTemplate struct {
	Files              []ProjectFile
	Directories        []string
	TemplatedVariables []string
}

type RenderedTemplate struct {
	Files       []ProjectFile
	Directories []string
}

func (r *RenderedTemplate) Equals(other *RenderedTemplate) bool {
	if len(r.Files) != len(other.Files) {
		return false
	}

	if len(r.Directories) != len(other.Directories) {
		return false
	}

	for i := range r.Files {
		if r.Files[i].Path != other.Files[i].Path {
			return false
		}

		if r.Files[i].Content != other.Files[i].Content {
			return false
		}
	}

	for i := range r.Directories {
		if r.Directories[i] != other.Directories[i] {
			return false
		}
	}

	return true
}

func renderString(input string, vars map[string]interface{}) string {
	tmpl, err := template.New("").Parse(input)
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

func renderTemplate(input ProjectTemplate, userVars map[string]interface{}) (RenderedTemplate, error) {

	for _, v := range input.TemplatedVariables {
		if _, ok := userVars[v]; ok {
			continue
		} else {
			return RenderedTemplate{}, errors.New("Missing variable definition")
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
