package interceptor

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/fahmyabdul/gits-assignments/test-3/pkg/common"
	"github.com/fahmyabdul/gits-assignments/test-3/pkg/logger"
)

// MetadataInterceptor: Intercept gRPC metadata
func MetadataInterceptor(l *logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("could not grab metadata from context")
		}
		l.Info().Interface("metadata", meta).Msg("request metadata...")

		// Authorization
		authorization := meta.Get("authorization")
		if len(authorization) == 0 {
			return nil, errors.New("unauthorized")
		}

		if !strings.Contains(authorization[0], "Bearer") {
			return nil, errors.New("unauthorized")
		}

		token := strings.Split(authorization[0], " ")[1]
		if token == "" {
			return nil, errors.New("unauthorized")
		}

		// The token should be a jwt token
		err := common.VerifyJwt(token)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
