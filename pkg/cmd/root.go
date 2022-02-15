package cmd

import (
	"github.com/adamdusty/projectgen/pkg/pgen"
	"github.com/spf13/cobra"
)

var config pgen.Config

var rootCmd = &cobra.Command{
	Use:   "pgen",
	Short: "Project generator",
	Long:  "Project generator for all kinds of projects and custom templates!",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	// cmd, _, err := rootCmd.Find(os.Args[1:])

	// // default cmd if no cmd is given
	// if err == nil && cmd.Use == rootCmd.Use && cmd.Flags().Parse(os.Args[1:]) != pflag.ErrHelp {
	// 	args := append([]string{"generate"}, os.Args[1:]...)
	// 	rootCmd.SetArgs(args)
	// }

	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	rootCmd.PersistentFlags().StringVar(&config.Path, "config", "", "config file (default: ~/.config/.pgen/config.yaml or %APPDATA%/.pgen/config.yaml)")
}
