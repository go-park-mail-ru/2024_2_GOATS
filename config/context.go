package config

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/logger"
)

type ContextConfigKey struct{}
type ContextLoggerKey struct{}

func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ContextConfigKey{}, cfg)
}

func WrapRedisContext(ctx context.Context, cfg *Redis) context.Context {
	return context.WithValue(ctx, ContextConfigKey{}, cfg)
}

func WrapLocalStorageContext(ctx context.Context, cfg *LocalStorage) context.Context {
	return context.WithValue(ctx, ContextConfigKey{}, cfg)
}

func WrapLoggerContext(ctx context.Context, lg *logger.BaseLogger) context.Context {
	return context.WithValue(ctx, ContextLoggerKey{}, lg)
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

func FromBaseContext(ctx context.Context) (*logger.BaseLogger, string) {
	requestId := ctx.Value("request-id").(string)
	lg, ok := ctx.Value(ContextLoggerKey{}).(*logger.BaseLogger)
	if !ok {
		return logger.NewLogger(), requestId
	}

	return lg, requestId
}
