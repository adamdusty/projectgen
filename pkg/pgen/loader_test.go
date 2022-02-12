package pgen

import (
	"testing"
)

var templateJson string = `
{
    "directories": [
        "src",
        "docs",
        "tests"
    ],
    "files": [
        {
            "path": "src/main.cpp",
            "content": "#include <iostream>\n#inlcude \"lib.hpp\"\n\nauto main(int argc, char *argv[]) -> int {\n    std::cout << {{ project_name | string.downcase }}::greeting() << '\\n';\n    return 0;\n}"
        },
        {
            "path": "src/lib.hpp",
            "content": "#pragma once\n\nnamespace {{ project_name | string.downcase }} {\n\nauto greeting() -> char *;\n\n}"
        },
        {
            "path": "src/lib.cpp",
            "content": "#include \"lib.hpp\"\n\nnamespace {{ project_name | lower }} {\n\nauto greeting() -> char * {\n    return \"Hello from {{ project_name }}!\";\n}\n\n}"
        },
        {
            "path": "README.md",
            "content": "# {{ project_name }}"
        }
    ],
    "variables": [
		{
            "identifier": "project_name",
            "representation": "Project Name"
        }
    ]
}
`

var templateYaml string = `---
directories:
- src
- docs
- tests
files:
- path: src/main.cpp
  content: |-
    #include <iostream>
    #inlcude "lib.hpp"

    auto main(int argc, char *argv[]) -> int {
        std::cout << {{ project_name | string.downcase }}::greeting() << '\n';
        return 0;
    }
- path: src/lib.hpp
  content: |-
    #pragma once

    namespace {{ project_name | string.downcase }} {

    auto greeting() -> char *;

    }
- path: src/lib.cpp
  content: |-
    #include "lib.hpp"

    namespace {{ project_name | lower }} {

    auto greeting() -> char * {
        return "Hello from {{ project_name }}!";
    }

    }
- path: README.md
  content: "# {{ project_name }}"
variables:
- identifier: project_name
  representation: Project Name
`

var expectedTemplate ProjectTemplate = ProjectTemplate{

	Directories: []string{"src", "docs", "tests"},
	Files: []ProjectFile{
		{Path: "src/main.cpp", Content: "#include <iostream>\n#inlcude \"lib.hpp\"\n\nauto main(int argc, char *argv[]) -> int {\n    std::cout << {{ project_name | string.downcase }}::greeting() << '\\n';\n    return 0;\n}"},
		{Path: "src/lib.hpp", Content: "#pragma once\n\nnamespace {{ project_name | string.downcase }} {\n\nauto greeting() -> char *;\n\n}"},
		{Path: "src/lib.cpp", Content: "#include \"lib.hpp\"\n\nnamespace {{ project_name | lower }} {\n\nauto greeting() -> char * {\n    return \"Hello from {{ project_name }}!\";\n}\n\n}"},
		{Path: "README.md", Content: "# {{ project_name }}"},
	},
	Variables: []TemplateVariable{
		{"project_name", "Project Name", "", "", ""},
	},
}

func TestLoadFromJson(t *testing.T) {
	actual, err := LoadFromJson([]byte(templateJson))

	if err != nil {
		t.Error(err)
	}

	if !actual.Equals(&expectedTemplate) {
		t.Error("Loaded template not equal to expected")
	}
}

func TestLoadFromYaml(t *testing.T) {
	actual, err := LoadFromYaml([]byte(templateYaml))
	if err != nil {
		t.Error(err)
	}
	// runtime.Breakpoint()

	if !actual.Equals(&expectedTemplate) {
		t.Error("Loaded template not equal to expected")
	}
}
