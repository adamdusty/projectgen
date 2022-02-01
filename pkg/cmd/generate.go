package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
	generateCmd.Flags().StringVar(&template, "template", "t", "Name/Alias of template to use.")

	rootCmd.AddCommand(generateCmd)
}

func generate(cmd *cobra.Command, args []string) {
	fmt.Println("Hello from generate")
	// find template based on template alias

	// load template from file

	// generate project
}
