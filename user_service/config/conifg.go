package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config contains user_service configuration
type Config struct {
	Listener  Listener  `yaml:"listener"`
	Databases Databases `yaml:"databases"`
}

// Listener contains user_service listener port
type Listener struct {
	Port string `yaml:"port"`
}

// Databases contains user_service databases configuration
type Databases struct {
	Postgres     Postgres     `yaml:"postgres"`
	LocalStorage LocalStorage `yaml:"localStorage"`
}

// Postgres contains user_service postgres configuration
type Postgres struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	Name            string `yaml:"name"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
	ConnMaxIdleTime int    `yaml:"connMaxIdleTime"`
}

// LocalStorage contains user_service local storage configuration
type LocalStorage struct {
	UserAvatarsFullURL     string `yaml:"userAvatarsFullURL"`
	UserAvatarsRelativeURL string `yaml:"userAvatarsRelativeURL"`
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
