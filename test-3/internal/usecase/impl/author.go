package impl

import (
	"context"
	"fmt"

	"github.com/fahmyabdul/gits-assignments/test-3/internal/entity"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/usecase"
	"github.com/fahmyabdul/gits-assignments/test-3/pkg/logger"
)

type UsecaseAuthorImpl struct {
	repo   usecase.RepoAuthor
	logger *logger.Logger
}

func NewUsecaseAuthorImpl(r usecase.RepoAuthor, l *logger.Logger) *UsecaseAuthorImpl {
	return &UsecaseAuthorImpl{
		repo:   r,
		logger: l,
	}
}

func (p *UsecaseAuthorImpl) Create(ctx context.Context, request *entity.Author) error {
	if request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	err := p.repo.Create(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (p *UsecaseAuthorImpl) FetchById(ctx context.Context, id int) (*entity.Author, error) {
	if id == 0 {
		return nil, fmt.Errorf(_emptyRequest)
	}

	result, err := p.repo.FetchById(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *UsecaseAuthorImpl) FetchByName(ctx context.Context, name string) ([]*entity.Author, error) {
	if name == "" {
		return nil, fmt.Errorf(_emptyRequest)
	}

	result, err := p.repo.FetchByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *UsecaseAuthorImpl) FetchAll(ctx context.Context) ([]*entity.Author, error) {
	result, err := p.repo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *UsecaseAuthorImpl) Update(ctx context.Context, id int, request *entity.Author) error {
	if id == 0 || request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	err := p.repo.Update(ctx, id, request)
	if err != nil {
		return err
	}

	return nil
}

func (p *UsecaseAuthorImpl) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return fmt.Errorf(_emptyRequest)
	}

	err := p.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
