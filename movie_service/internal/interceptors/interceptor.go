package interceptors

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/metrics"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ctxKey int

const (
	requestIDKey ctxKey = iota
	service             = "movie_microservice"
)

func WithLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Info().Msg("No metadata found in context")
	}

	reqID := md.Get("request_id")
	if len(reqID) > 0 {
		ctx = context.WithValue(ctx, requestIDKey, reqID[0])
		logger := log.With().Str("request_id", reqID[0]).Caller().Logger()
		ctx = logger.WithContext(ctx)
	} else {
		log.Info().Msg("Empty request_id in metadata")
	}

	return handler(ctx, req)
}

func AccessLogInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()
	logger := log.Ctx(ctx)

	defer func() {
		if r := recover(); r != nil {
			duration := time.Since(start)

			logger.Info().
				Str("method", info.FullMethod).
				Interface("request_body", req).
				Dur("work_time", duration).
				Interface("panic", r).
				Msg("panic occurred")

			metrics.GRPCServerRequestsTotal.WithLabelValues(service, info.FullMethod, "panic").Inc()
			metrics.GRPCServerDuration.WithLabelValues(service, info.FullMethod).Observe(duration.Seconds())

			err = status.Errorf(codes.Internal, "panic occurred: %v", r)
		}
	}()

	resp, err = handler(ctx, req)

	duration := time.Since(start)
	statusCode := codes.OK
	if err != nil {
		statusCode = status.Code(err)
	}

	metrics.GRPCServerRequestsTotal.WithLabelValues(service, info.FullMethod, statusCode.String()).Inc()
	metrics.GRPCServerDuration.WithLabelValues(service, info.FullMethod).Observe(duration.Seconds())

	var acMsg string
	if err != nil {
		acMsg = "Request completed with error"
	} else {
		acMsg = "Request completed successfully"
	}

	logger.Info().
		Str("method", info.FullMethod).
		Interface("request_body", req).
		Dur("work_time", duration).
		Interface("response_body", resp).
		Msg(acMsg)

	return resp, err
}
