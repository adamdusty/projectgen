package cmd

import (
	"io"
	"strings"
	"testing"

	"github.com/adamdusty/projectgen/pkg/pgen"
)

func TestFindTemplate(t *testing.T) { t.Error("unimpl") }

func TestLoadTemplateFile(t *testing.T) { t.Error("unimpl") }

func TestQueryUserVars(t *testing.T) {}

func TestQueryVar(t *testing.T) {
	v := pgen.TemplateVariable{
		Default:          "test",
		Identifier:       "identifier",
		Representation:   "Identifier",
		ShortDescription: "Test identifier",
		LongDescription:  "This is a test variable and the long description is unused for now",
	}

	expected := "unit_test_proj_name"
	actual, err := queryVar(&v, strings.NewReader("unit_test_proj_name"), io.Discard)

	if actual != expected {
		t.Errorf("Query return did not match expected: %s != %s", actual, expected)
	}

	if err != nil {
		t.Errorf("Produced unexpected error: %s", err)
	}
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

	if actual != expected {
		t.Errorf("Processed input did not match expected: %s != %s", actual, expected)
	}

	if err != nil {
		t.Errorf("Produced unexpected error: %s", err)
	}
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

	if actual != expected {
		t.Errorf("Generated prompt not equal to expected: %s", actual)
	}
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

	if actual != expected {
		t.Errorf("Generated prompt not equal to expected: %s", actual)
	}
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

	if actual != expected {
		t.Errorf("Generated prompt not equal to expected: %s", actual)
	}
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

	if actual != expected {
		t.Errorf("Generated prompt not equal to expected: %s", actual)
	}
}
