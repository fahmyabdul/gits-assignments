package usecase

import (
	"context"

	"github.com/fahmyabdul/gits-assignments/test-3/internal/entity"
)

type UsecaseAuthor interface {
	Create(ctx context.Context, request *entity.Author) error
	FetchById(ctx context.Context, id int) (*entity.Author, error)
	FetchByName(ctx context.Context, name string) ([]*entity.Author, error)
	FetchAll(ctx context.Context) ([]*entity.Author, error)
	Update(ctx context.Context, id int, request *entity.Author) error
	Delete(ctx context.Context, id int) error
}

type UsecasePublisher interface {
	Create(ctx context.Context, request *entity.Publisher) error
	FetchById(ctx context.Context, id int) (*entity.Publisher, error)
	FetchByName(ctx context.Context, name string) ([]*entity.Publisher, error)
	FetchAll(ctx context.Context) ([]*entity.Publisher, error)
	Update(ctx context.Context, id int, request *entity.Publisher) error
	Delete(ctx context.Context, id int) error
}

type UsecaseBook interface {
	Create(ctx context.Context, request *entity.Book, author_id int) error
	FetchById(ctx context.Context, id int) (*entity.Book, error)
	FetchByName(ctx context.Context, name string) ([]*entity.Book, error)
	FetchAll(ctx context.Context) ([]*entity.Book, error)
	Update(ctx context.Context, id, author_id int, request *entity.Book) error
	Delete(ctx context.Context, id int) error
	FetchByAuthorId(ctx context.Context, author_id int) (*entity.BookByAuthor, error)
}

type RepoAuthor interface {
	Create(ctx context.Context, request *entity.Author) error
	FetchById(ctx context.Context, id int) (*entity.Author, error)
	FetchByName(ctx context.Context, name string) ([]*entity.Author, error)
	FetchAll(ctx context.Context) ([]*entity.Author, error)
	Update(ctx context.Context, id int, request *entity.Author) error
	Delete(ctx context.Context, id int) error
}

type RepoPublisher interface {
	Create(ctx context.Context, request *entity.Publisher) error
	FetchById(ctx context.Context, id int) (*entity.Publisher, error)
	FetchByName(ctx context.Context, name string) ([]*entity.Publisher, error)
	FetchAll(ctx context.Context) ([]*entity.Publisher, error)
	Update(ctx context.Context, id int, request *entity.Publisher) error
	Delete(ctx context.Context, id int) error
}

type RepoBook interface {
	Create(ctx context.Context, request *entity.Book, author_id int) error
	FetchById(ctx context.Context, id int) (*entity.Book, error)
	FetchByName(ctx context.Context, name string) ([]*entity.Book, error)
	FetchAll(ctx context.Context) ([]*entity.Book, error)
	Update(ctx context.Context, id, author_id int, request *entity.Book) error
	Delete(ctx context.Context, id int) error
	FetchByAuthorId(ctx context.Context, author_id int) (*entity.BookByAuthor, error)
}
