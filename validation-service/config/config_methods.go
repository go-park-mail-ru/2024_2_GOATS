package config

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func GetConfigFromContext(ctx context.Context) *Config {
	value, ok := ctx.Value(ConfigContextKey{}).(*Config)
	if !ok {
		return nil
	}

	return value
}

func CreateConfigContext() (context.Context, error) {
	ctx := context.Background()
	config, err := readConf()
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, ConfigContextKey{}, config), nil
}

func readConf() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}

	cfg := &Config{}

	filename, _ := filepath.Abs(os.Getenv("CFG_PATH"))
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
