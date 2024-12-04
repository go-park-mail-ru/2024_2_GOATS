package config

import "context"

type ConfigContextKey struct{}
type LocalStorageContextKey struct{}
type CurrentUserContextKey struct{}

func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ConfigContextKey{}, cfg)
}

func FromContext(ctx context.Context) *Config {
	cfg, ok := ctx.Value(ConfigContextKey{}).(*Config)
	if !ok {
		return nil
	}

	return cfg
}
