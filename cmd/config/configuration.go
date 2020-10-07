package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is a global object that holds configuration.
var Config Configuration

// Configuration contains server and database configurations.
type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// LoadConfiguration reads the yaml file and get the desired configuration.
func LoadConfiguration() error {
	viper.SetConfigName("cozy")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	return viper.Unmarshal(&Config)
}
