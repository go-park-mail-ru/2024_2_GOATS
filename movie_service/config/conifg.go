package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Listener  Listener  `yaml:"listener"`
	Databases Databases `yaml:"databases"`
}

type Databases struct {
	Postgres     Postgres     `yaml:"postgres"`
	Redis        Redis        `yaml:"redis"`
	LocalStorage LocalStorage `yaml:"localStorage"`
}

type LocalStorage struct {
	UserAvatarsFullURL     string `yaml:"userAvatarsFullURL"`
	UserAvatarsRelativeURL string `yaml:"userAvatarsRelativeURL"`
}

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

type Redis struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Cookie Cookie `yaml:"cookie"`
}

type Cookie struct {
	Name   string        `yaml:"name"`
	MaxAge time.Duration `yaml:"maxAge"`
}

type Listener struct {
	Address     string        `yaml:"address"`
	Port        string        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

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
	viper.AddConfigPath("..")

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
