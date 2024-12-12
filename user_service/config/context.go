package config

import "context"

// ContextConfigKey is a context key for full config
type ContextConfigKey struct{}

// LocalStorageContextKey is a context key for local storage config
type LocalStorageContextKey struct{}

// WrapContext wraps full config into context
func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ContextConfigKey{}, cfg)
}

// WrapLocalStorageContext wraps local storage config into context
func WrapLocalStorageContext(ctx context.Context, locSt *LocalStorage) context.Context {
	return context.WithValue(ctx, LocalStorageContextKey{}, locSt)
}

// FromContext gets full config from context
func FromContext(ctx context.Context) *Config {
	cfg, ok := ctx.Value(ContextConfigKey{}).(*Config)
	if !ok {
		return nil
	}

	return cfg
}

// FromLocalStorageContext gets local storage config from context
func FromLocalStorageContext(ctx context.Context) *LocalStorage {
	locST, ok := ctx.Value(LocalStorageContextKey{}).(*LocalStorage)
	if !ok {
		return nil
	}

	return locST
}
