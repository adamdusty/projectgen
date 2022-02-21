package cmd

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"github.com/adamdusty/projectgen/pkg/pgen"
	"github.com/stretchr/testify/assert"
)

func TestFindTemplate(t *testing.T) {
	// Not sure how to test this function
	// May need refactor
}

func TestLoadTemplateFile(t *testing.T) {
	// Not sure how to test this function
	// May need refactor
}

// Issue with this test
// queryVar function creates new scanner every time a definition is requested,
func TestQueryUserVars(t *testing.T) {
	vars := []pgen.TemplateVariable{
		{
			Default:          "noname",
			Identifier:       "proj_name",
			Representation:   "Project Name",
			ShortDescription: "Name of project",
			LongDescription:  "The name you wish your project to be generated with",
		},
		{
			Default:          "me",
			Identifier:       "author",
			Representation:   "Author",
			ShortDescription: "Author of project",
			LongDescription:  "The person responsible for writing the project",
		},
	}

	expected := map[string]interface{}{
		"proj_name": "TestProjName",
		"author":    "JohnDoe",
	}

	input := strings.NewReader("TestProjName\nJohnDoe\n")
	actual := queryUserVars(vars, input, io.Discard)

	assert.Equal(t, expected, actual, "actual user defs don't match expected")
}

func TestQueryVar(t *testing.T) {
	v := pgen.TemplateVariable{
		Default:          "test",
		Identifier:       "identifier",
		Representation:   "Identifier",
		ShortDescription: "Test identifier",
		LongDescription:  "This is a test variable and the long description is unused for now",
	}

	input := strings.NewReader("unit_test_proj_name")
	scanner := bufio.NewScanner(input)

	expected := "unit_test_proj_name"
	actual := queryVar(buildQueryPrompt(&v), scanner, io.Discard)

	assert.Equal(t, expected, actual, "actual response doesn't match expected")
}

func TestBuildQueryPromptAllFieldsPresent(t *testing.T) {
	data := pgen.TemplateVariable{
		Default:          "test",
		Identifier:       "identifier",
		Representation:   "Identifier",
		ShortDescription: "Test identifier",
		LongDescription:  "This is a test variable and the long description is unused for now",
	}

	expected := "Identifier (Test identifier) [test]: "
	actual := buildQueryPrompt(&data)

	assert.Equal(t, expected, actual, "actual prompt doesn't match expected", expected, actual)
}

func TestBuildQueryPromptNoDefault(t *testing.T) {
	data := pgen.TemplateVariable{
		Default:          "",
		Identifier:       "identifier",
		Representation:   "Identifier",
		ShortDescription: "Test identifier",
		LongDescription:  "This is a test variable and the long description is unused for now",
	}

	expected := "Identifier (Test identifier): "
	actual := buildQueryPrompt(&data)

	assert.Equal(t, expected, actual, "actual prompt doesn't match expected", expected, actual)
}

func TestBuildQueryPromptNoDefaultNoDescription(t *testing.T) {
	data := pgen.TemplateVariable{
		Default:          "",
		Identifier:       "identifier",
		Representation:   "Identifier",
		ShortDescription: "",
		LongDescription:  "This is a test variable and the long description is unused for now",
	}

	expected := "Identifier: "
	actual := buildQueryPrompt(&data)

	assert.Equal(t, expected, actual, "actual prompt doesn't match expected", expected, actual)
}

func TestBuildQueryPromptNoDescription(t *testing.T) {
	data := pgen.TemplateVariable{
		Default:          "test",
		Identifier:       "identifier",
		Representation:   "Identifier",
		ShortDescription: "",
		LongDescription:  "This is a test variable and the long description is unused for now",
	}

	expected := "Identifier [test]: "
	actual := buildQueryPrompt(&data)

	assert.Equal(t, expected, actual, "actual prompt doesn't match expected", expected, actual)
}
