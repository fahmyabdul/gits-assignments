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

type GrpcBookHandler struct {
	usecase *impl.UsecaseBookImpl
	logger  *logger.Logger

	proto.UnsafeBookServiceServer
}

func NewGrpcBookHandler(usecase *impl.UsecaseBookImpl, logger *logger.Logger) *GrpcBookHandler {
	return &GrpcBookHandler{
		usecase: usecase,
		logger:  logger,
	}
}

func (p *GrpcBookHandler) protoToEntity(protoData *proto.Book) (*entity.Book, error) {
	if protoData == nil {
		return nil, fmt.Errorf("proto data must not empty")
	}

	output := &entity.Book{
		Id:          protoData.Id,
		Name:        protoData.Name,
		Pages:       protoData.Pages,
		PublisherId: protoData.PublisherId,
		CreatedAt:   protoData.CreatedAt.AsTime(),
		UpdatedAt:   protoData.UpdatedAt.AsTime(),
	}

	return output, nil
}

func (p *GrpcBookHandler) entityToProto(entityData *entity.Book) (*proto.Book, error) {
	if entityData == nil {
		return nil, fmt.Errorf("entity data must not empty")
	}

	output := &proto.Book{
		Id:          entityData.Id,
		Name:        entityData.Name,
		Pages:       entityData.Pages,
		PublisherId: entityData.PublisherId,
		CreatedAt:   timestamppb.New(entityData.CreatedAt),
		UpdatedAt:   timestamppb.New(entityData.UpdatedAt),
	}

	return output, nil
}

func (p *GrpcBookHandler) CreateBook(ctx context.Context, in *proto.BookCreateRequest) (*proto.BookCreateResponse, error) {
	if in == nil {
		return &proto.BookCreateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	inputData, err := p.protoToEntity(in.GetRequest())
	if err != nil {
		return &proto.BookCreateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	err = p.usecase.Create(ctx, inputData, int(in.GetAuthorId()))
	if err != nil {
		return &proto.BookCreateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.BookCreateResponse{
		Status: "success",
	}, nil
}

func (p *GrpcBookHandler) FetchByIdBook(ctx context.Context, in *proto.BookFetchByIdRequest) (*proto.BookFetchByIdResponse, error) {
	if in == nil {
		return &proto.BookFetchByIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	fetchedData, err := p.usecase.FetchById(ctx, int(in.GetId()))
	if err != nil {
		return &proto.BookFetchByIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	protoData, err := p.entityToProto(fetchedData)
	if err != nil {
		return &proto.BookFetchByIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.BookFetchByIdResponse{
		Status: "success",
		Data:   protoData,
	}, nil
}

func (p *GrpcBookHandler) FetchByNameBook(ctx context.Context, in *proto.BookFetchByNameRequest) (*proto.BookFetchByNameResponse, error) {
	if in == nil {
		return &proto.BookFetchByNameResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	fetchedData, err := p.usecase.FetchByName(ctx, in.GetName())
	if err != nil {
		return &proto.BookFetchByNameResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	var outputProtoData []*proto.Book

	for _, entityData := range fetchedData {
		protoData, err := p.entityToProto(entityData)
		if err != nil {
			return &proto.BookFetchByNameResponse{
					Status: "failed",
					Data:   nil,
				}, status.Errorf(
					codes.Internal,
					err.Error(),
				)
		}

		outputProtoData = append(outputProtoData, protoData)
	}

	return &proto.BookFetchByNameResponse{
		Status: "success",
		Data:   outputProtoData,
	}, nil
}

func (p *GrpcBookHandler) FetchAllBook(ctx context.Context, in *emptypb.Empty) (*proto.BookFetchAllResponse, error) {
	if in == nil {
		return &proto.BookFetchAllResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	fetchedData, err := p.usecase.FetchAll(ctx)
	if err != nil {
		return &proto.BookFetchAllResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	var outputProtoData []*proto.Book

	for _, entityData := range fetchedData {
		protoData, err := p.entityToProto(entityData)
		if err != nil {
			return &proto.BookFetchAllResponse{
					Status: "failed",
					Data:   nil,
				}, status.Errorf(
					codes.Internal,
					err.Error(),
				)
		}

		outputProtoData = append(outputProtoData, protoData)
	}

	return &proto.BookFetchAllResponse{
		Status: "success",
		Data:   outputProtoData,
	}, nil
}

func (p *GrpcBookHandler) UpdateBook(ctx context.Context, in *proto.BookUpdateRequest) (*proto.BookUpdateResponse, error) {
	if in == nil {
		return &proto.BookUpdateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	inputData, err := p.protoToEntity(in.GetRequest())
	if err != nil {
		return &proto.BookUpdateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	err = p.usecase.Update(ctx, int(in.GetId()), int(in.GetAuthorId()), inputData)
	if err != nil {
		return &proto.BookUpdateResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.BookUpdateResponse{
		Status: "success",
	}, nil
}

func (p *GrpcBookHandler) DeleteBook(ctx context.Context, in *proto.BookDeleteRequest) (*proto.BookDeleteResponse, error) {
	if in == nil {
		return &proto.BookDeleteResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	err := p.usecase.Delete(ctx, int(in.GetId()))
	if err != nil {
		return &proto.BookDeleteResponse{
				Status: "failed",
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	return &proto.BookDeleteResponse{
		Status: "success",
	}, nil
}

func (p *GrpcBookHandler) FetchByAuthorId(ctx context.Context, in *proto.BookFetchByAuthorIdRequest) (*proto.BookFetchByAuthorIdResponse, error) {
	if in == nil {
		return &proto.BookFetchByAuthorIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				_emptyRequest,
			)
	}

	fetchedData, err := p.usecase.FetchByAuthorId(ctx, int(in.GetAuthorId()))
	if err != nil {
		return &proto.BookFetchByAuthorIdResponse{
				Status: "failed",
				Data:   nil,
			}, status.Errorf(
				codes.Internal,
				err.Error(),
			)
	}

	var outputProtoData []*proto.Book

	for _, entityData := range fetchedData.Books {
		protoData, err := p.entityToProto(&entityData)
		if err != nil {
			return &proto.BookFetchByAuthorIdResponse{
					Status: "failed",
					Data:   nil,
				}, status.Errorf(
					codes.Internal,
					err.Error(),
				)
		}

		outputProtoData = append(outputProtoData, protoData)
	}

	return &proto.BookFetchByAuthorIdResponse{
		Status:     "success",
		AuthorId:   fetchedData.AuthorId,
		AuthorName: fetchedData.AuthorName,
		Data:       outputProtoData,
	}, nil
}
