package config

import "time"

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

func NewConfig() *Config {
	return &Config{}
}
