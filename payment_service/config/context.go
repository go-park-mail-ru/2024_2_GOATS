package config

import "context"

// ContextConfigKey is a context key for full config
type ContextConfigKey struct{}

// WrapContext wraps full config into context
func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ContextConfigKey{}, cfg)
}

// FromContext gets full config from context
func FromContext(ctx context.Context) *Config {
	cfg, ok := ctx.Value(ContextConfigKey{}).(*Config)
	if !ok {
		return nil
	}

	return cfg
}
