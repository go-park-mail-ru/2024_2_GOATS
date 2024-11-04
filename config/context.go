package config

import (
	"context"
)

type ContextConfigKey struct{}

func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ContextConfigKey{}, cfg)
}

func WrapRedisContext(ctx context.Context, cfg *Redis) context.Context {
	return context.WithValue(ctx, ContextConfigKey{}, cfg)
}

func WrapLocalStorageContext(ctx context.Context, cfg *LocalStorage) context.Context {
	return context.WithValue(ctx, ContextConfigKey{}, cfg)
}

func FromContext(ctx context.Context) *Config {
	value, ok := ctx.Value(ContextConfigKey{}).(*Config)

	if !ok {
		return nil
	}

	return value
}

func FromRedisContext(ctx context.Context) *Redis {
	value, ok := ctx.Value(ContextConfigKey{}).(*Redis)

	if !ok {
		return nil
	}

	return value
}

func FromLocalStorageContext(ctx context.Context) *LocalStorage {
	value, ok := ctx.Value(ContextConfigKey{}).(*LocalStorage)
	if !ok {
		return nil
	}

	return value
}
