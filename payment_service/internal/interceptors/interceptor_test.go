package interceptors

import (
	"context"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestWithLogger(t *testing.T) {
	t.Run("adds logger and request_id to context when metadata is present", func(t *testing.T) {
		md := metadata.New(map[string]string{"request_id": "12345"})
		ctx := metadata.NewIncomingContext(context.Background(), md)

		handler := func(ctx context.Context, _ interface{}) (interface{}, error) {
			reqID := ctx.Value(requestIDKey).(string)
			assert.Equal(t, "12345", reqID)

			logger := zerolog.Ctx(ctx)
			assert.NotNil(t, logger)

			return "success", nil
		}

		resp, err := WithLogger(ctx, nil, nil, handler)
		assert.NoError(t, err)
		assert.Equal(t, "success", resp)
	})

	t.Run("logs missing metadata when no request_id is present", func(t *testing.T) {
		ctx := context.Background()

		handler := func(ctx context.Context, _ interface{}) (interface{}, error) {
			reqID := ctx.Value(requestIDKey)
			assert.Nil(t, reqID)

			return "success", nil
		}

		resp, err := WithLogger(ctx, nil, nil, handler)
		assert.NoError(t, err)
		assert.Equal(t, "success", resp)
	})
}

func TestAccessLogInterceptor_NoMetrics(t *testing.T) {
	t.Run("logs successful request", func(t *testing.T) {
		handler := func(_ context.Context, _ interface{}) (interface{}, error) {
			return "response", nil
		}

		info := &grpc.UnaryServerInfo{FullMethod: "/test/method"}
		start := time.Now()
		resp, err := AccessLogInterceptor(context.Background(), "request", info, handler)
		duration := time.Since(start)

		assert.NoError(t, err)
		assert.Equal(t, "response", resp)

		assert.LessOrEqual(t, duration.Seconds(), 5.0, "Request took too long")
	})

	t.Run("logs error request", func(t *testing.T) {
		handler := func(_ context.Context, _ interface{}) (interface{}, error) {
			return nil, status.Error(codes.Internal, "internal error")
		}

		info := &grpc.UnaryServerInfo{FullMethod: "/test/method"}
		resp, err := AccessLogInterceptor(context.Background(), "request", info, handler)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("handles panic and logs it", func(t *testing.T) {
		handler := func(_ context.Context, _ interface{}) (interface{}, error) {
			panic("panic error")
		}

		info := &grpc.UnaryServerInfo{FullMethod: "/test/method"}
		resp, err := AccessLogInterceptor(context.Background(), "request", info, handler)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Contains(t, err.Error(), "panic occurred")
	})
}
