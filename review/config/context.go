package config

import (
	"context"
)

type ContextConfigKey struct{}
type ContextRedisKey struct{}
type ContextLocalStorageKey struct{}

func WrapRedisContext(ctx context.Context, cfg *Redis) context.Context {
	return context.WithValue(ctx, ContextRedisKey{}, cfg)
}

func FromRedisContext(ctx context.Context) *Redis {
	value, ok := ctx.Value(ContextRedisKey{}).(*Redis)

	if !ok {
		return nil
	}

	return value
}

type ConfigContextKey struct{}

func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ConfigContextKey{}, cfg)
}
