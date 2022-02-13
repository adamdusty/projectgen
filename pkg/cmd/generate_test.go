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

	actual, err := queryUserVars(vars, strings.NewReader("TestProjName\nJohnDoe"), io.Discard)

	assert.Nil(t, err, "unexpected err:", err)
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

	expected := "unit_test_proj_name"
	actual, err := queryVar(&v, bufio.NewScanner(input), io.Discard)

	assert.Nil(t, err, "unexpected err:", err)
	assert.Equal(t, expected, actual, "actual response doesn't match expected")
}

func TestProcessInput(t *testing.T) {
	v := pgen.TemplateVariable{
		Default:          "test",
		Identifier:       "identifier",
		Representation:   "Identifier",
		ShortDescription: "Test identifier",
		LongDescription:  "This is a test variable and the long description is unused for now",
	}

	expected := "unit_test_proj_name"
	actual, err := processInput(&v, "unit_test_proj_name")

	assert.Nil(t, err, "unexpected err:", err)
	assert.Equalf(t, expected, actual, "Processed input did not match expected: %s != %s", actual, expected)
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
