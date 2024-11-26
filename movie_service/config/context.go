package config

import "context"

type ConfigContextKey struct{}
type LocalStorageContextKey struct{}
type CurrentUserContextKey struct{}

func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ConfigContextKey{}, cfg)
}

func WrapLocalStorageContext(ctx context.Context, locSt *LocalStorage) context.Context {
	return context.WithValue(ctx, LocalStorageContextKey{}, locSt)
}

func FromContext(ctx context.Context) *Config {
	cfg, ok := ctx.Value(ConfigContextKey{}).(*Config)
	if !ok {
		return nil
	}

	return cfg
}

func FromLocalStorageContext(ctx context.Context) *LocalStorage {
	locST, ok := ctx.Value(LocalStorageContextKey{}).(*LocalStorage)
	if !ok {
		return nil
	}

	return locST
}

type ContextConfigKey struct{}
type ContextRedisKey struct{}
type ContextLocalStorageKey struct{}
type CurrentUserKey struct{}

func FromRedisContext(ctx context.Context) *Redis {
	value, ok := ctx.Value(ContextRedisKey{}).(*Redis)

	if !ok {
		return nil
	}

	return value
}

func CurrentUserID(ctx context.Context) int {
	value, ok := ctx.Value(CurrentUserKey{}).(int)
	if !ok {
		return 0
	}

	return value
}
