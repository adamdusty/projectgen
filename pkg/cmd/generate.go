package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var outputDir string
var template string

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
	fmt.Println("Hello from generate")

	// - Find specified template in template directory
	// - Search through template directory for file base/cpp-exe ($HOME/.pgen/templates/base/cpp-exe)
	// file, err := findTemplate(template)
	// if err != nil {
	// 	panic(err)
	// }

	// - Load template from serialization format

	// - Generate project at path given as first argument

}

func findTemplate(alias string) (*os.File, error) {
	// If path is absolute, try to use file at absolute path
	file, err := os.Open(alias)
	if err == nil {
		return file, nil
	}

	// If path is relative, search for file at CWD + alias
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := filepath.Join(cwd, alias)

	file, err = os.Open(path)
	if err == nil {
		return file, nil
	}

	// If not found, try to read template directory from config, and attempt to resolve from there
	dir := viper.GetString("template-directory")
	path = filepath.Join(dir, alias)

	file, err = os.Open(path)
	if err == nil {
		return file, nil
	}

	return nil, errors.New("Unable to find template: " + alias)
}
