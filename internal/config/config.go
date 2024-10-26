package config

import (
	"github.com/spf13/viper"
)

// Config holds all configuration for our program
type Config struct {
	ServerAddress string `mapstructure:"server_address"`
	ModelPath     string `mapstructure:"model_path"`
}

// Load loads configuration from environment variables and config file
func Load() (*Config, error) {
	viper.SetDefault("server_address", ":8080")
	viper.SetDefault("model_path", "./models/model.joblib")

	viper.AutomaticEnv()

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
