package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fahmyabdul/gits-assignments/test-3/internal/entity"

	"github.com/fahmyabdul/gits-assignments/test-3/pkg/pggorm"
	"github.com/fahmyabdul/gits-assignments/test-3/pkg/redispkg"
)

type RepoBookImpl struct {
	pg   *pggorm.PgGorm
	rpkg *redispkg.RedisPkg
}

// NewRepoBookImpl : Initialize new Book Repository
func NewRepoBookImpl(pg *pggorm.PgGorm, rpkg *redispkg.RedisPkg) *RepoBookImpl {
	return &RepoBookImpl{
		pg:   pg,
		rpkg: rpkg,
	}
}

func (p *RepoBookImpl) updateBookAuthor(book_id int32, author_id int) error {
	inputData := entity.BookAuthor{
		BookId:   book_id,
		AuthorId: int32(author_id),
	}
	if p.pg.Conn.Model(&inputData).Where("book_id = ?", book_id).Updates(&inputData).RowsAffected == 0 {
		// If not found then Create new Book Author
		fmt.Println("Book author with given book_id not found, create new instead of update")
		createResult := p.pg.Conn.Create(inputData)
		if createResult.Error != nil {
			return createResult.Error
		}
	}

	return nil
}

// Create :
func (p *RepoBookImpl) Create(ctx context.Context, request *entity.Book, author_id int) error {
	if request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	var countExist int32
	p.pg.Conn.Model(&entity.Book{}).Where("name = ?", request.Name).Count(&countExist)
	if countExist > 0 {
		return fmt.Errorf("book with same name already exists")
	}

	var outputAuthor entity.Author
	authorResult := p.pg.Conn.Model(outputAuthor).Where("id = ?", author_id).Find(&outputAuthor)
	if authorResult.Error != nil {
		return authorResult.Error
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	result := p.pg.Conn.Create(request)
	if result.Error != nil {
		return result.Error
	}

	err := p.updateBookAuthor(request.Id, author_id)
	if err != nil {
		return err
	}

	jsonByte, err := json.Marshal(request)
	if err != nil {
		return err
	}

	_, err = p.rpkg.Pool.Get().Do("HSET", request.KeyName(), request.Id, string(jsonByte))
	if err != nil {
		return err
	}

	return nil
}

// FetchById :
func (p *RepoBookImpl) FetchById(ctx context.Context, id int) (*entity.Book, error) {
	if id == 0 {
		return nil, fmt.Errorf(_emptyRequest)
	}

	var output entity.Book

	result := p.pg.Conn.Where("id = ?", id).Find(&output)
	if result.Error != nil {
		return nil, result.Error
	}

	return &output, nil
}

// FetchByName :
func (p *RepoBookImpl) FetchByName(ctx context.Context, name string) ([]*entity.Book, error) {
	if name == "" {
		return nil, fmt.Errorf(_emptyRequest)
	}

	var output []*entity.Book

	result := p.pg.Conn.Where("name LIKE ?", "%"+name+"%").Find(&output)
	if result.Error != nil {
		return nil, result.Error
	}

	return output, nil
}

// FetchAll :
func (p *RepoBookImpl) FetchAll(ctx context.Context) ([]*entity.Book, error) {
	var outputs []*entity.Book

	result := p.pg.Conn.Find(&outputs)
	if result.Error != nil {
		return nil, result.Error
	}

	return outputs, nil
}

// Update :
func (p *RepoBookImpl) Update(ctx context.Context, id, author_id int, request *entity.Book) error {
	if id == 0 || request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	var countExist int32
	p.pg.Conn.Model(&entity.Book{}).Where("id = ?", id).Count(&countExist)
	if countExist == 0 {
		return fmt.Errorf("book with given id not exists")
	}

	request.CreatedAt = time.Time{}
	request.UpdatedAt = time.Now()

	result := p.pg.Conn.Model(&entity.Book{}).Where("id = ?", id).Updates(request)
	if result.Error != nil {
		return result.Error
	}

	err := p.updateBookAuthor(int32(id), author_id)
	if err != nil {
		return err
	}

	jsonByte, err := json.Marshal(request)
	if err != nil {
		return err
	}

	_, err = p.rpkg.Pool.Get().Do("HSET", request.KeyName(), request.Id, string(jsonByte))
	if err != nil {
		return err
	}

	return nil
}

// Delete :
func (p *RepoBookImpl) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return fmt.Errorf(_emptyRequest)
	}

	var countExist int32
	p.pg.Conn.Model(&entity.Book{}).Where("id = ?", id).Count(&countExist)
	if countExist == 0 {
		return fmt.Errorf("book with given id not exists")
	}

	result := p.pg.Conn.Where("id = ?", id).Delete(&entity.Book{})
	if result.Error != nil {
		return result.Error
	}

	bookAuthorResult := p.pg.Conn.Where("book_id = ?", id).Delete(&entity.BookAuthor{})
	if bookAuthorResult.Error != nil {
		return bookAuthorResult.Error
	}

	_, err := p.rpkg.Pool.Get().Do("HDEL", (&entity.Book{}).KeyName(), id)
	if err != nil {
		return err
	}

	return nil
}

// FetchByAuthorId :
func (p *RepoBookImpl) FetchByAuthorId(ctx context.Context, author_id int) (*entity.BookByAuthor, error) {
	if author_id == 0 {
		return nil, fmt.Errorf(_emptyRequest)
	}

	var joinOutputs []*entity.BookByAuthorTmp

	result := p.pg.Conn.Table("t_book_author").
		Select("t_book_author.author_id, t_author.name as author_name, t_book.*").
		Joins("left join t_author on t_author.id = t_book_author.author_id").
		Joins("left join t_book on t_book.id = t_book_author.book_id ").
		Where("t_book_author.author_id = ?", author_id).
		Find(&joinOutputs)
	if result.Error != nil {
		return nil, result.Error
	}

	var output entity.BookByAuthor

	for _, joinData := range joinOutputs {
		output.AuthorId = joinData.AuthorId
		output.AuthorName = joinData.AuthorName
		output.Books = append(output.Books, entity.Book{
			Id:          joinData.Id,
			Name:        joinData.Name,
			Pages:       joinData.Pages,
			PublisherId: joinData.PublisherId,
			CreatedAt:   joinData.CreatedAt,
			UpdatedAt:   joinData.UpdatedAt,
		})
	}

	return &output, nil
}
