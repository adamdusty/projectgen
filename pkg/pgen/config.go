package pgen

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Path              string // Path to config file
	TemplateDirectory string // Path to search for templates
}

func (c *Config) InitConfig() {
	// TODO: Set up proper logging
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	// Set default config file path if not provided
	if c.Path == "" {
		configDir, err := os.UserConfigDir()
		if err != nil {
			panic(err)
		}

		c.Path = filepath.Join(configDir, ".pgen", "config.yaml")
	}

	file, err := os.OpenFile(c.Path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	if err := viper.ReadConfig(file); err == nil {
		log.Infof("Using config found at: %s", c.Path)
	}
}

func defaultConfig() *Config {
	cache, err := os.UserCacheDir()
	if err != nil {
		log.Panic(err)
	}

	templateDir := filepath.Join(cache, ".pgen", "templates")

	return &Config{
		Path:              "",
		TemplateDirectory: templateDir,
	}
}
