package entity

import (
	"strings"
	"time"
)

type Book struct {
	Id          int32     `json:"id,omitempty" validate:"required"`
	Name        string    `json:"name,omitempty" validate:"required"`
	Pages       int32     `json:"pages,omitempty" validate:"required"`
	PublisherId int32     `json:"publisher_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (p *Book) TableName() string {
	return "t_book"
}

func (p *Book) KeyName() string {
	return strings.Replace(p.TableName(), "_", ":", -1)
}

type BookAuthor struct {
	BookId   int32 `json:"book_id" validate:"required"`
	AuthorId int32 `json:"author_id" validate:"required"`
}

func (p *BookAuthor) TableName() string {
	return "t_book_author"
}

type BookByAuthorTmp struct {
	AuthorId    int32     `json:"author_id" validate:"required"`
	AuthorName  string    `json:"author_name,omitempty" validate:"required"`
	Id          int32     `json:"id,omitempty" validate:"required"`
	Name        string    `json:"name,omitempty" validate:"required"`
	Pages       int32     `json:"pages,omitempty" validate:"required"`
	PublisherId int32     `json:"publisher_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type BookByAuthor struct {
	AuthorId   int32  `json:"author_id" validate:"required"`
	AuthorName string `json:"author_name,omitempty" validate:"required"`
	Books      []Book `json:"books,omitempty" validate:"required"`
}
