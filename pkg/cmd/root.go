package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "project-gen",
	Short:   "Project generator",
	Long:    "Project generator for all kinds of projects and custom templates!",
	Aliases: []string{"pgen"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from pgen cmd")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
