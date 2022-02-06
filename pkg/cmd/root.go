package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgPath     string
	templateDir string
)

var rootCmd = &cobra.Command{
	Use:     "project-gen",
	Short:   "Project generator",
	Long:    "Project generator for all kinds of projects and custom templates!",
	Aliases: []string{"pgen"},
	Run:     func(cmd *cobra.Command, args []string) {},
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", "", "config file (default: $HOME/.pgen)")
	rootCmd.PersistentFlags().StringVar(&templateDir, "template-directory", "", "directory with template definitions (default: $HOME/.pgen/templates")

	viper.BindPFlag("template-directory", rootCmd.PersistentFlags().Lookup("template-directory"))

}

func initConfig() {
	if cfgPath != "" {

		viper.SetConfigFile(cfgPath)
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file: ", viper.ConfigFileUsed())
		}

	}
	// else {
	// 	home, err := os.UserHomeDir()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if _, err := os.Stat(filepath.Join(home, ".pgen")); os.IsNotExist(err) {
	// 		err := os.MkdirAll(filepath.Join(home, ".pgen"), 0755)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}

	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigType("yaml")
	// 	viper.SetConfigName(".pgen")
	// }
}
