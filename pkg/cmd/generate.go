package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/adamdusty/projectgen/pkg/pgen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var outputDir string
var template string

type UserInputError struct {
	Input string
}

var generateCmd = &cobra.Command{
	Use:   "generate ",
	Short: "Generate project.",
	Args:  cobra.MinimumNArgs(1),
	Run:   generate,
}

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	generateCmd.Flags().StringVarP(&outputDir, "output", "o", cwd, "Directory to generate project in. Defaults to current working directory.")
	generateCmd.Flags().StringVar(&template, "template", "", "Name/Alias of template to use.")

	rootCmd.AddCommand(generateCmd)
}

func generate(cmd *cobra.Command, args []string) {
	path := args[0]

	// - Find specified template in template directory
	// - Search through template directory for file (base/cpp-exe = $HOME/.pgen/templates/base/cpp-exe)
	projectTemplate, err := loadTemplateFile(path)
	if err != nil {
		panic(err)
	}

	// Query user for variable defs
	userDefs, err := queryUserVars(projectTemplate.Variables, os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}

	// Render template strings
	renderedTemplate, err := pgen.RenderTemplate(projectTemplate, userDefs)
	if err != nil {
		panic(err)
	}

	// - Generate project at path given as first argument
	if !filepath.IsAbs(path) {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		path = filepath.Join(cwd, path)
	}

	// Generate project
	err = pgen.GenerateProject(path, renderedTemplate)
	if err != nil {
		panic(err)
	}
}

func findTemplate(alias string) (*os.File, error) {
	// If path is absolute, try to use file at absolute path
	file, err := os.Open(alias)
	if err == nil {
		return file, err
	}

	// If path is relative, search for file at CWD + alias
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := filepath.Join(cwd, alias)

	file, err = os.Open(path)
	if err == nil {
		return file, err
	}

	// If not found, try to read template directory from config, and attempt to resolve from there
	dir := viper.GetString("template-directory")
	path = filepath.Join(dir, alias)

	file, err = os.Open(path)
	if err == nil {
		return file, err
	}

	return nil, errors.New("Unable to find template: " + alias)
}

func loadTemplateFile(path string) (*pgen.ProjectTemplate, error) {
	file, err := findTemplate(template)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	tmpl, err := pgen.LoadFromYaml(buf.Bytes())
	if err != nil {
		return tmpl, err
	}

	return tmpl, nil
}

func queryUserVars(vars []pgen.TemplateVariable, reader io.Reader, writer io.Writer) (map[string]interface{}, error) {
	defs := make(map[string]interface{})

	for i := 0; i < len(vars); {
		def, err := queryVar(&vars[i], os.Stdin, os.Stdout)

		if err != nil {
			switch e := err.(type) {
			case *UserInputError:
				output := fmt.Sprintf("%s is required.", vars[i].Representation)
				writer.Write([]byte(output))
				continue
			default:
				return nil, e
			}
		}

		defs[vars[i].Identifier] = def
		i++
	}

	return defs, nil
}

func queryVar(v *pgen.TemplateVariable, reader io.Reader, writer io.Writer) (string, error) {
	writer.Write([]byte(buildQueryPrompt(v)))
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	def, err := processInput(v, scanner.Text())
	return def, err
}

func processInput(v *pgen.TemplateVariable, input string) (string, error) {
	if v.Default != "" && input == "" {
		return v.Default, nil
	}

	if v.Default == "" && input == "" {
		return "", &UserInputError{input}
	}

	return input, nil
}

func buildQueryPrompt(v *pgen.TemplateVariable) string {
	var prompt strings.Builder

	prompt.WriteString(v.Representation)

	if v.ShortDescription != "" {
		prompt.WriteRune(' ')
		prompt.WriteRune('(')
		prompt.WriteString(v.ShortDescription)
		prompt.WriteRune(')')
	}

	if v.Default != "" {
		prompt.WriteRune(' ')
		prompt.WriteRune('[')
		prompt.WriteString(v.Default)
		prompt.WriteRune(']')
	}

	prompt.WriteString(": ")

	return prompt.String()
}

func (e *UserInputError) Error() string {
	if e.Input == "" {
		return "user input was empty"
	}

	return fmt.Sprintf("user input error: %s", e.Input)
}
