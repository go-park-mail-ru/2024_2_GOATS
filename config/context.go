package config

import (
	"context"
)

// ContextConfigKey is a context key for full config
type ContextConfigKey struct{}

// CurrentUserKey is a context key for current user id
type CurrentUserKey struct{}

// WrapContext wraps full config into context
func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ContextConfigKey{}, cfg)
}

// FromContext gets full config from context
func FromContext(ctx context.Context) *Config {
	value, ok := ctx.Value(ContextConfigKey{}).(*Config)

	if !ok {
		return nil
	}

	return value
}

// CurrentUserID gets current user id from context
func CurrentUserID(ctx context.Context) int {
	value, ok := ctx.Value(CurrentUserKey{}).(int)
	if !ok {
		return 0
	}

	return value
}
