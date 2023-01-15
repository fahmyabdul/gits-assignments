package interceptor

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/fahmyabdul/gits-assignments/test-3/pkg/logger"
)

func MetadataInterceptor(l *logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("could not grab metadata from context")
		}
		l.Info().Interface("metadata", meta).Msg("metadata from client")

		return handler(ctx, req)
	}
}
