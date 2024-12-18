package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config contains facade configuration
type Config struct {
	Listener  Listener  `yaml:"listener"`
	Databases Databases `yaml:"databases"`
}

// Databases contains auth_service databases configuration
type Databases struct {
	Redis Redis `yaml:"redis"`
}

// Redis contains auth_service redis configuration
type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Listener contains facade listener params
type Listener struct {
	Address     string        `yaml:"address"`
	Port        int           `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

// New returns an instance of Config
func New(isTest bool) (*Config, error) {
	err := setupViper(isTest)
	if err != nil {
		return nil, fmt.Errorf("config creation error: %w", err)
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the config file: %w", err)
	}

	return cfg, nil
}

func setupViper(isTest bool) error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read .env file: %v", err)
	}

	var cfgName string
	if isTest {
		cfgName = "config_test"
	} else {
		cfgName = "config"
	}

	viper.SetConfigName(cfgName)
	viper.SetConfigType("yml")
	viper.AddConfigPath(viper.GetString("VIPER_CFG_PATH"))

	err = viper.MergeInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config.yml file: %v", err)
	}

	return nil
}
