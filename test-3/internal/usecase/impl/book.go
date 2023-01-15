package impl

import (
	"context"
	"fmt"

	"github.com/fahmyabdul/gits-assignments/test-3/internal/entity"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/usecase"
	"github.com/fahmyabdul/gits-assignments/test-3/pkg/logger"
)

type UsecaseBookImpl struct {
	repo   usecase.RepoBook
	logger *logger.Logger
}

func NewUsecaseBookImpl(r usecase.RepoBook, l *logger.Logger) *UsecaseBookImpl {
	return &UsecaseBookImpl{
		repo:   r,
		logger: l,
	}
}

func (p *UsecaseBookImpl) Create(ctx context.Context, request *entity.Book, author_id int) error {
	if request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	err := p.repo.Create(ctx, request, author_id)
	if err != nil {
		return err
	}

	return nil
}

func (p *UsecaseBookImpl) FetchById(ctx context.Context, id int) (*entity.Book, error) {
	if id == 0 {
		return nil, fmt.Errorf(_emptyRequest)
	}

	result, err := p.repo.FetchById(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *UsecaseBookImpl) FetchByName(ctx context.Context, name string) ([]*entity.Book, error) {
	if name == "" {
		return nil, fmt.Errorf(_emptyRequest)
	}

	result, err := p.repo.FetchByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *UsecaseBookImpl) FetchAll(ctx context.Context) ([]*entity.Book, error) {
	result, err := p.repo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *UsecaseBookImpl) Update(ctx context.Context, id, author_id int, request *entity.Book) error {
	if id == 0 || request == nil || author_id == 0 {
		return fmt.Errorf(_emptyRequest)
	}

	err := p.repo.Update(ctx, id, author_id, request)
	if err != nil {
		return err
	}

	return nil
}

func (p *UsecaseBookImpl) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return fmt.Errorf(_emptyRequest)
	}

	err := p.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *UsecaseBookImpl) FetchByAuthorId(ctx context.Context, author_id int) (*entity.BookByAuthor, error) {
	if author_id == 0 {
		return nil, fmt.Errorf(_emptyRequest)
	}

	result, err := p.repo.FetchByAuthorId(ctx, author_id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
