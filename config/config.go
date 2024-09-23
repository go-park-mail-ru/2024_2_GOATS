package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Listener Listener `yaml:"listener"`
	Postgres Postgres `yaml:"postgres"`
}

type ConfigContextKey struct{}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Listener struct {
	Address     string        `yaml:"address"`
	Port        int           `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func New() (*Config, error) {
	err := setupViper()
	if err != nil {
		return nil, fmt.Errorf("config creation error: %w", err)
	}

	listner := Listener{
		Address:     viper.GetString("listener.address"),
		Port:        viper.GetInt("listener.port"),
		Timeout:     viper.GetDuration("listener.timeout"),
		IdleTimeout: viper.GetDuration("listener.idle_timeout"),
	}

	postgres := Postgres{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetInt("postgres.port"),
		User:     viper.GetString("postgres.user"),
		Password: viper.GetString("postgres.password"),
		Name:     viper.GetString("postgres.name"),
	}

	return &Config{
		Listener: listner,
		Postgres: postgres,
	}, nil
}

func setupViper() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read .env file: %v", err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(viper.GetString("VIPER_CFG_PATH"))

	err = viper.MergeInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config.yml file: %v", err)
	}

	return nil
}
