package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/adamdusty/projectgen/pkg/pgen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var template string
var templateDir string

var generateCmd = &cobra.Command{
	Use:   "generate [output dir]",
	Short: "Generate project.",
	Args:  cobra.MinimumNArgs(1),

	Run: generate,
}

func init() {
	generateCmd.Flags().StringVar(&template, "template", "", "path or name of template")
	generateCmd.Flags().StringVar(&templateDir, "template-directory", "", "directory with templates")

	generateCmd.MarkFlagRequired("template")

	// viper.BindPFlag("template", generateCmd.Flags().Lookup("template"))
	viper.BindPFlag("template-directory", generateCmd.Flags().Lookup("template-directory"))

	rootCmd.AddCommand(generateCmd)
}

func generate(cmd *cobra.Command, args []string) {
	path := args[0]

	projectTemplate, err := loadTemplateFile(template)
	if err != nil {
		panic(err)
	}

	// Query user for variable defs
	userDefs := queryUserVars(projectTemplate.Variables, os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}

	// Generate project
	err = pgen.GenerateProject(path, projectTemplate, userDefs)
	if err != nil {
		panic(err)
	}
}

func findTemplate(alias string) (*os.File, error) {
	// If template directory is set, check directory for alias
	if viper.IsSet("template-directory") {
		dir := viper.GetString("template-directory")
		path := filepath.Join(dir, alias)

		file, err := os.Open(path)
		return file, err
	} else {
		file, err := os.Open(alias)
		return file, err
	}
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

func queryVar(prompt string, scanner *bufio.Scanner, writer io.Writer) string {
	writer.Write([]byte(prompt))
	scanner.Scan()
	return scanner.Text()
}

func queryUserVars(vars []pgen.TemplateVariable, reader io.Reader, writer io.Writer) map[string]interface{} {

	defs := make(map[string]interface{})
	scanner := bufio.NewScanner(reader)

	for i := 0; i < len(vars); {
		v := &vars[i]
		def := queryVar(buildQueryPrompt(v), scanner, writer)

		if def == "" && v.Default == "" {
			fmt.Fprintf(writer, "[ERROR] %s is required.\n", v.Representation)
			continue
		}

		if def == "" {
			def = v.Default
		}

		defs[v.Identifier] = def
		i++
	}

	return defs
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
