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

type GrpcAuthorHandler struct {
	usecase *impl.UsecaseAuthorImpl
	logger  *logger.Logger

	proto.UnsafeAuthorServiceServer
}

func NewGrpcAuthorHandler(usecase *impl.UsecaseAuthorImpl, logger *logger.Logger) *GrpcAuthorHandler {
	return &GrpcAuthorHandler{
		usecase: usecase,
		logger:  logger,
	}
}

func (p *GrpcAuthorHandler) protoToEntity(protoData *proto.Author) (*entity.Author, error) {
	if protoData == nil {
		return nil, fmt.Errorf("proto data must not empty")
	}

	output := &entity.Author{
		Id:        protoData.Id,
		Name:      protoData.Name,
		Detail:    protoData.Detail,
		CreatedAt: protoData.CreatedAt.AsTime(),
		UpdatedAt: protoData.UpdatedAt.AsTime(),
	}

	return output, nil
}

func (p *GrpcAuthorHandler) entityToProto(entityData *entity.Author) (*proto.Author, error) {
	if entityData == nil {
		return nil, fmt.Errorf("entity data must not empty")
	}

	output := &proto.Author{
		Id:        entityData.Id,
		Name:      entityData.Name,
		Detail:    entityData.Detail,
		CreatedAt: timestamppb.New(entityData.CreatedAt),
		UpdatedAt: timestamppb.New(entityData.UpdatedAt),
	}

	return output, nil
}

func (p *GrpcAuthorHandler) CreateAuthor(ctx context.Context, in *proto.AuthorCreateRequest) (*proto.AuthorCreateResponse, error) {
	if in == nil {
		return &proto.AuthorCreateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	inputData, err := p.protoToEntity(in.GetRequest())
	if err != nil {
		return &proto.AuthorCreateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	err = p.usecase.Create(ctx, inputData)
	if err != nil {
		return &proto.AuthorCreateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.AuthorCreateResponse{
		Status: "success",
	}, nil
}

func (p *GrpcAuthorHandler) FetchByIdAuthor(ctx context.Context, in *proto.AuthorFetchByIdRequest) (*proto.AuthorFetchByIdResponse, error) {
	if in == nil {
		return &proto.AuthorFetchByIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	fetchedData, err := p.usecase.FetchById(ctx, int(in.GetId()))
	if err != nil {
		return &proto.AuthorFetchByIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	protoData, err := p.entityToProto(fetchedData)
	if err != nil {
		return &proto.AuthorFetchByIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.AuthorFetchByIdResponse{
		Status: "success",
		Data:   protoData,
	}, nil
}

func (p *GrpcAuthorHandler) FetchByNameAuthor(ctx context.Context, in *proto.AuthorFetchByNameRequest) (*proto.AuthorFetchByNameResponse, error) {
	if in == nil {
		return &proto.AuthorFetchByNameResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	fetchedData, err := p.usecase.FetchByName(ctx, in.GetName())
	if err != nil {
		return &proto.AuthorFetchByNameResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	var outputProtoData []*proto.Author

	for _, entityData := range fetchedData {
		protoData, err := p.entityToProto(entityData)
		if err != nil {
			return &proto.AuthorFetchByNameResponse{
					Status: "failed",
					Data:   nil,
				}, status.Errorf(
					codes.Internal,
					err.Error(),
				)
		}

		outputProtoData = append(outputProtoData, protoData)
	}

	return &proto.AuthorFetchByNameResponse{
		Status: "success",
		Data:   outputProtoData,
	}, nil
}

func (p *GrpcAuthorHandler) FetchAllAuthor(ctx context.Context, in *emptypb.Empty) (*proto.AuthorFetchAllResponse, error) {
	if in == nil {
		return &proto.AuthorFetchAllResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	fetchedData, err := p.usecase.FetchAll(ctx)
	if err != nil {
		return &proto.AuthorFetchAllResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	var outputProtoData []*proto.Author

	for _, entityData := range fetchedData {
		protoData, err := p.entityToProto(entityData)
		if err != nil {
			return &proto.AuthorFetchAllResponse{
					Status: "failed",
					Data:   nil,
				}, status.Errorf(
					codes.Internal,
					err.Error(),
				)
		}

		outputProtoData = append(outputProtoData, protoData)
	}

	return &proto.AuthorFetchAllResponse{
		Status: "success",
		Data:   outputProtoData,
	}, nil
}

func (p *GrpcAuthorHandler) UpdateAuthor(ctx context.Context, in *proto.AuthorUpdateRequest) (*proto.AuthorUpdateResponse, error) {
	if in == nil {
		return &proto.AuthorUpdateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	inputData, err := p.protoToEntity(in.GetRequest())
	if err != nil {
		return &proto.AuthorUpdateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	err = p.usecase.Update(ctx, int(in.GetId()), inputData)
	if err != nil {
		return &proto.AuthorUpdateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.AuthorUpdateResponse{
		Status: "success",
	}, nil
}

func (p *GrpcAuthorHandler) DeleteAuthor(ctx context.Context, in *proto.AuthorDeleteRequest) (*proto.AuthorDeleteResponse, error) {
	if in == nil {
		return &proto.AuthorDeleteResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	err := p.usecase.Delete(ctx, int(in.GetId()))
	if err != nil {
		return &proto.AuthorDeleteResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.AuthorDeleteResponse{
		Status: "success",
	}, nil
}
