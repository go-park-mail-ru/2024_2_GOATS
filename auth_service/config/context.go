package config

import (
	"context"
)

// ContextRedisKey is a context key for redis config
type ContextRedisKey struct{}

// WrapRedisContext wraps redis config into context
func WrapRedisContext(ctx context.Context, cfg *Redis) context.Context {
	return context.WithValue(ctx, ContextRedisKey{}, cfg)
}

// FromRedisContext gets redis config from context
func FromRedisContext(ctx context.Context) *Redis {
	value, ok := ctx.Value(ContextRedisKey{}).(*Redis)

	if !ok {
		return nil
	}

	return value
}
