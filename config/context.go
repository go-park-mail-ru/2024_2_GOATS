package config

import (
	"context"
)

func WrapContext(cfg *Config) (context.Context, error) {
	return context.WithValue(context.Background(), ConfigContextKey{}, cfg), nil
}

func FromContext(ctx context.Context) *Config {
	value, ok := ctx.Value(ConfigContextKey{}).(*Config)

	if !ok {
		return nil
	}

	return value
}
