package interceptor

import (
	"context"

	"google.golang.org/grpc"

	"github.com/fahmyabdul/gits-assignments/test-3/pkg/logger"
)

func LoggingInterceptor(l *logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		return handler(ctx, req)
	}
}
