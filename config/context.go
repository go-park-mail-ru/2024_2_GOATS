package config

import (
	"context"
)

type ContextConfigKey struct{}

func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ContextConfigKey{}, cfg)
}

func WrapRedisContext(ctx context.Context, cfg *Redis) context.Context {
	return context.WithValue(ctx, ConfigContextKey{}, cfg)
}

func FromContext(ctx context.Context) *Config {
	value, ok := ctx.Value(ContextConfigKey{}).(*Config)

	if !ok {
		return nil
	}

	return value
}
