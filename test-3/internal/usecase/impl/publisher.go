package impl

import (
	"context"
	"fmt"

	"github.com/fahmyabdul/gits-assignments/test-3/internal/entity"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/usecase"
	"github.com/fahmyabdul/gits-assignments/test-3/pkg/logger"
)

type UsecasePublisherImpl struct {
	repo   usecase.RepoPublisher
	logger *logger.Logger
}

func NewUsecasePublisherImpl(r usecase.RepoPublisher, l *logger.Logger) *UsecasePublisherImpl {
	return &UsecasePublisherImpl{
		repo:   r,
		logger: l,
	}
}

func (p *UsecasePublisherImpl) Create(ctx context.Context, request *entity.Publisher) error {
	if request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	err := p.repo.Create(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (p *UsecasePublisherImpl) FetchById(ctx context.Context, id int) (*entity.Publisher, error) {
	if id == 0 {
		return nil, fmt.Errorf(_emptyRequest)
	}

	result, err := p.repo.FetchById(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *UsecasePublisherImpl) FetchByName(ctx context.Context, name string) ([]*entity.Publisher, error) {
	if name == "" {
		return nil, fmt.Errorf(_emptyRequest)
	}

	result, err := p.repo.FetchByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *UsecasePublisherImpl) FetchAll(ctx context.Context) ([]*entity.Publisher, error) {
	result, err := p.repo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *UsecasePublisherImpl) Update(ctx context.Context, id int, request *entity.Publisher) error {
	if id == 0 || request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	err := p.repo.Update(ctx, id, request)
	if err != nil {
		return err
	}

	return nil
}

func (p *UsecasePublisherImpl) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return fmt.Errorf(_emptyRequest)
	}

	err := p.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
