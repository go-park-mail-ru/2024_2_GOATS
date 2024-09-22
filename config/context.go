package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func WrapContext(cfg *Config) (context.Context, error) {
	filename, _ := filepath.Abs(viper.GetString("CFG_PATH"))
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("failed to open env: %w", err)
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshall config.yml: %w", err)
	}

	return context.WithValue(context.Background(), ConfigContextKey{}, cfg), nil
}

func FromContext(ctx context.Context) *Config {
	value, ok := ctx.Value(ConfigContextKey{}).(*Config)

	if !ok {
		return nil
	}

	return value
}
