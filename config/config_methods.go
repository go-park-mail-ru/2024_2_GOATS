package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func CreateConfigContext() (context.Context, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	filename, _ := filepath.Abs(os.Getenv("CFG_PATH"))
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("failed to open env: %w", err)
	}

	cfg := &Config{}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshall config.yml: %w", err)
	}

	return context.WithValue(context.Background(), ConfigContextKey{}, cfg), nil
}

func GetConfigFromContext(ctx context.Context) *Config {
	value, ok := ctx.Value(ConfigContextKey{}).(*Config)

	if !ok {
		return nil
	}

	return value
}
