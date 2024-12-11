package interceptors

import (
	"context"

	"google.golang.org/grpc"
)

// ChainUnaryInterceptors creates a chain of interceptors
func ChainUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		chainedHandler := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			current := interceptors[i]
			chainedHandler = wrapHandler(current, chainedHandler, info)
		}

		return chainedHandler(ctx, req)
	}
}

func wrapHandler(
	interceptor grpc.UnaryServerInterceptor,
	handler grpc.UnaryHandler,
	info *grpc.UnaryServerInfo,
) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return interceptor(ctx, req, info, handler)
	}
}
