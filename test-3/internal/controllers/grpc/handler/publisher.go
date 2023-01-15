package handler

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/fahmyabdul/gits-assignments/test-3/pkg/logger"

	"github.com/fahmyabdul/gits-assignments/test-3/internal/entity"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/proto"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/usecase/impl"
)

type GrpcPublisherHandler struct {
	usecase *impl.UsecasePublisherImpl
	logger  *logger.Logger

	proto.UnsafePublisherServiceServer
}

func NewGrpcPublisherHandler(usecase *impl.UsecasePublisherImpl, logger *logger.Logger) *GrpcPublisherHandler {
	return &GrpcPublisherHandler{
		usecase: usecase,
		logger:  logger,
	}
}

func (p *GrpcPublisherHandler) protoToEntity(protoData *proto.Publisher) (*entity.Publisher, error) {
	if protoData == nil {
		return nil, fmt.Errorf("proto data must not empty")
	}

	output := &entity.Publisher{
		Id:        protoData.Id,
		Name:      protoData.Name,
		Detail:    protoData.Detail,
		CreatedAt: protoData.CreatedAt.AsTime(),
		UpdatedAt: protoData.UpdatedAt.AsTime(),
	}

	return output, nil
}

func (p *GrpcPublisherHandler) entityToProto(entityData *entity.Publisher) (*proto.Publisher, error) {
	if entityData == nil {
		return nil, fmt.Errorf("entity data must not empty")
	}

	output := &proto.Publisher{
		Id:        entityData.Id,
		Name:      entityData.Name,
		Detail:    entityData.Detail,
		CreatedAt: timestamppb.New(entityData.CreatedAt),
		UpdatedAt: timestamppb.New(entityData.UpdatedAt),
	}

	return output, nil
}

func (p *GrpcPublisherHandler) CreatePublisher(ctx context.Context, in *proto.PublisherCreateRequest) (*proto.PublisherCreateResponse, error) {
	if in == nil {
		return &proto.PublisherCreateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	inputData, err := p.protoToEntity(in.GetRequest())
	if err != nil {
		return &proto.PublisherCreateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	err = p.usecase.Create(ctx, inputData)
	if err != nil {
		return &proto.PublisherCreateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.PublisherCreateResponse{
		Status: "success",
	}, nil
}

func (p *GrpcPublisherHandler) FetchByIdPublisher(ctx context.Context, in *proto.PublisherFetchByIdRequest) (*proto.PublisherFetchByIdResponse, error) {
	if in == nil {
		return &proto.PublisherFetchByIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	fetchedData, err := p.usecase.FetchById(ctx, int(in.GetId()))
	if err != nil {
		return &proto.PublisherFetchByIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	protoData, err := p.entityToProto(fetchedData)
	if err != nil {
		return &proto.PublisherFetchByIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.PublisherFetchByIdResponse{
		Status: "success",
		Data:   protoData,
	}, nil
}

func (p *GrpcPublisherHandler) FetchByNamePublisher(ctx context.Context, in *proto.PublisherFetchByNameRequest) (*proto.PublisherFetchByNameResponse, error) {
	if in == nil {
		return &proto.PublisherFetchByNameResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	fetchedData, err := p.usecase.FetchByName(ctx, in.GetName())
	if err != nil {
		return &proto.PublisherFetchByNameResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	var outputProtoData []*proto.Publisher

	for _, entityData := range fetchedData {
		protoData, err := p.entityToProto(entityData)
		if err != nil {
			return &proto.PublisherFetchByNameResponse{
					Status: "failed",
					Data:   nil,
				}, status.Errorf(
					codes.Internal,
					err.Error(),
				)
		}

		outputProtoData = append(outputProtoData, protoData)
	}

	return &proto.PublisherFetchByNameResponse{
		Status: "success",
		Data:   outputProtoData,
	}, nil
}

func (p *GrpcPublisherHandler) FetchAllPublisher(ctx context.Context, in *emptypb.Empty) (*proto.PublisherFetchAllResponse, error) {
	if in == nil {
		return &proto.PublisherFetchAllResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	fetchedData, err := p.usecase.FetchAll(ctx)
	if err != nil {
		return &proto.PublisherFetchAllResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	var outputProtoData []*proto.Publisher

	for _, entityData := range fetchedData {
		protoData, err := p.entityToProto(entityData)
		if err != nil {
			return &proto.PublisherFetchAllResponse{
					Status: "failed",
					Data:   nil,
				}, status.Errorf(
					codes.Internal,
					err.Error(),
				)
		}

		outputProtoData = append(outputProtoData, protoData)
	}

	return &proto.PublisherFetchAllResponse{
		Status: "success",
		Data:   outputProtoData,
	}, nil
}

func (p *GrpcPublisherHandler) UpdatePublisher(ctx context.Context, in *proto.PublisherUpdateRequest) (*proto.PublisherUpdateResponse, error) {
	if in == nil {
		return &proto.PublisherUpdateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	inputData, err := p.protoToEntity(in.GetRequest())
	if err != nil {
		return &proto.PublisherUpdateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	err = p.usecase.Update(ctx, int(in.GetId()), inputData)
	if err != nil {
		return &proto.PublisherUpdateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.PublisherUpdateResponse{
		Status: "success",
	}, nil
}

func (p *GrpcPublisherHandler) DeletePublisher(ctx context.Context, in *proto.PublisherDeleteRequest) (*proto.PublisherDeleteResponse, error) {
	if in == nil {
		return &proto.PublisherDeleteResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	err := p.usecase.Delete(ctx, int(in.GetId()))
	if err != nil {
		return &proto.PublisherDeleteResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.PublisherDeleteResponse{
		Status: "success",
	}, nil
}
