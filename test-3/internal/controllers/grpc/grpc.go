package grpc

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/fahmyabdul/gits-assignments/test-3/internal/controllers/grpc/handler"

	"github.com/fahmyabdul/gits-assignments/test-3/config"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/controllers/grpc/interceptor"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/proto"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/usecase/impl"
	"github.com/fahmyabdul/gits-assignments/test-3/pkg/logger"
)

type GrpcCtrl struct {
	usecase map[string]interface{}
}

func NewGrpcCtrl(logger *logger.Logger, config *config.GrpcConfig, opts ...Option) error {
	grpcCtrl := &GrpcCtrl{
		usecase: make(map[string]interface{}),
	}

	for _, opt := range opts {
		opt(grpcCtrl)
	}

	listen, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		return err
	}

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		interceptor.LoggingInterceptor(logger),
		interceptor.MetadataInterceptor(logger),
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
	)
	reflection.Register(grpcServer)

	authorHandler := handler.NewGrpcAuthorHandler(grpcCtrl.usecase["author"].(*impl.UsecaseAuthorImpl), logger)
	publisherHandler := handler.NewGrpcPublisherHandler(grpcCtrl.usecase["publisher"].(*impl.UsecasePublisherImpl), logger)
	bookHandler := handler.NewGrpcBookHandler(grpcCtrl.usecase["book"].(*impl.UsecaseBookImpl), logger)
	proto.RegisterAuthorServiceServer(grpcServer, authorHandler)
	proto.RegisterPublisherServiceServer(grpcServer, publisherHandler)
	proto.RegisterBookServiceServer(grpcServer, bookHandler)

	logger.Info().Str("port", ":"+config.Port).Msg("Starting Grpc Server")
	if err := grpcServer.Serve(listen); err != nil {
		return err
	}

	return nil
}
