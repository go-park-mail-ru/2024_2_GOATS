package config

import (
	"context"
)

type ConfigContextKey struct{}

func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ConfigContextKey{}, cfg)
}

func WrapRedisContext(ctx context.Context, cfg *Redis) context.Context {
	return context.WithValue(ctx, ConfigContextKey{}, cfg)
}

func FromContext(ctx context.Context) *Config {
	value, ok := ctx.Value(ConfigContextKey{}).(*Config)

	if !ok {
		return nil
	}

	return value
}

func FromRedisContext(ctx context.Context) *Redis {
	value, ok := ctx.Value(ConfigContextKey{}).(*Redis)

	if !ok {
		return nil
	}

	return value
}
