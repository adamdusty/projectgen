package cmd

import (
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
	templateDir := viper.GetString("template-directory")
	templatePath := filepath.Join(templateDir, cmd.Flags().Lookup("template").Value.String())

	// - Search through template directory for file base/cpp-exe ($HOME/.pgen/templates/base/cpp-exe)
	// - Load template from serialization format
	// - Generate project at path given as first argument

}
