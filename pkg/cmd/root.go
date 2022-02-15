package cmd

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgPath string

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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", "", "config file (default: ~/.config/.pgen/config.yaml or %APPDATA%/.pgen/config.yaml)")
}

func initConfig() {
	if cfgPath != "" {
		viper.SetConfigFile(cfgPath)
	} else {
		viper.SetConfigName("config")

		searchPath, err := configSearchPath()
		if err != nil {
			panic(err)
		}

		viper.AddConfigPath(searchPath)
	}

	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		switch e := err.(type) {
		case viper.ConfigFileNotFoundError:
			promptForConfigCreation(os.Stdin, os.Stdout)
		default:
			panic(e)
		}
	}
}

func configSearchPath() (string, error) {
	path, err := os.UserConfigDir()
	if err == nil {
		path = filepath.Join(path, ".pgen")
	}

	return path, err
}

func promptForConfigCreation(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)

	writer.Write([]byte("Config not found, would you like to generate one? [Y/n]: "))
	in.Scan()
	response := strings.ToLower(in.Text())

	for response != "y" && response != "n" && response != "" {
		writer.Write([]byte("Invalid response. Would you like to generate a config? [Y/n]: "))
		in.Scan()
		response = strings.ToLower(in.Text())
	}

	if response == "y" || response == "" {
		err := createConfigFile()
		if err != nil {
			panic(err)
		}
	} else {
		return
	}
}

func createConfigFile() error {
	viper.WriteConfig()
	return nil
}
