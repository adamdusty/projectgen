package template

import "testing"

func TestRenderStringRendersExpectedValues(t *testing.T) {
	content := "{{ .proj_name }}"
	vars := make(map[string]interface{})
	vars["proj_name"] = "HelloProj"

	actual := renderString(content, vars)
	expected := "HelloProj"

	if actual != expected {
		t.Errorf("Actual: %s, Expected: %s", actual, expected)
	}
}

func TestRenderTemplateHappyPath(t *testing.T) {
	files := make([]ProjectFile, 3)
	files[0] = ProjectFile{Path: "src/main.cpp", Content: "Hello {{ .proj_name }}!"}
	files[1] = ProjectFile{Path: "src/lib.hpp", Content: "Hello {{ .proj_name }} lib!"}
	files[2] = ProjectFile{Path: "src/lib.cpp", Content: "Hello {{ .proj_name }} impl!"}

	dirs := make([]string, 3)
	dirs[0] = "docs"
	dirs[1] = "src"
	dirs[2] = "build"

	tv := make([]string, 1)
	tv[0] = "proj_name"

	tmpl := ProjectTemplate{Files: files, Directories: dirs, TemplatedVariables: tv}

	vars := make(map[string]interface{})
	vars["proj_name"] = "TestProj"

	actual, _ := renderTemplate(tmpl, vars)
	expected := RenderedTemplate{
		Files: []ProjectFile{
			{Path: "src/main.cpp", Content: "Hello TestProj!"},
			{Path: "src/lib.hpp", Content: "Hello TestProj lib!"},
			{Path: "src/lib.cpp", Content: "Hello TestProj impl!"},
		},
		Directories: []string{
			"docs",
			"src",
			"build",
		},
	}

	if !actual.Equals(&expected) {
		t.Error("Render templates are not equivalent")
	}
}
