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

type RepoAuthorImpl struct {
	pg   *pggorm.PgGorm
	rpkg *redispkg.RedisPkg
}

// NewRepoAuthorImpl : Initialize new Author Repository
func NewRepoAuthorImpl(pg *pggorm.PgGorm, rpkg *redispkg.RedisPkg) *RepoAuthorImpl {
	return &RepoAuthorImpl{
		pg:   pg,
		rpkg: rpkg,
	}
}

// Create :
func (p *RepoAuthorImpl) Create(ctx context.Context, request *entity.Author) error {
	if request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	// var countExist int32
	// p.pg.Conn.Model(&entity.Author{}).Where("name = ?", request.Name).Count(&countExist)
	// if countExist > 0 {
	// 	return fmt.Errorf("author with same name already exists")
	// }

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	result := p.pg.Conn.Create(request)
	if result.Error != nil {
		return result.Error
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
func (p *RepoAuthorImpl) FetchById(ctx context.Context, id int) (*entity.Author, error) {
	if id == 0 {
		return nil, fmt.Errorf(_emptyRequest)
	}

	var output entity.Author

	result := p.pg.Conn.Where("id = ?", id).Find(&output)
	if result.Error != nil {
		return nil, result.Error
	}

	return &output, nil
}

// FetchByName :
func (p *RepoAuthorImpl) FetchByName(ctx context.Context, name string) ([]*entity.Author, error) {
	if name == "" {
		return nil, fmt.Errorf(_emptyRequest)
	}

	var output []*entity.Author

	result := p.pg.Conn.Where("name LIKE ?", "%"+name+"%").Find(&output)
	if result.Error != nil {
		return nil, result.Error
	}

	return output, nil
}

// FetchAll :
func (p *RepoAuthorImpl) FetchAll(ctx context.Context) ([]*entity.Author, error) {
	var outputs []*entity.Author

	result := p.pg.Conn.Find(&outputs)
	if result.Error != nil {
		return nil, result.Error
	}

	return outputs, nil
}

// Update :
func (p *RepoAuthorImpl) Update(ctx context.Context, id int, request *entity.Author) error {
	if id == 0 || request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	var countExist int32
	p.pg.Conn.Model(&entity.Author{}).Where("id = ?", id).Count(&countExist)
	if countExist == 0 {
		return fmt.Errorf("author with given id not exists")
	}

	request.CreatedAt = time.Time{}
	request.UpdatedAt = time.Now()

	result := p.pg.Conn.Model(&entity.Author{}).Where("id = ?", id).Updates(request)
	if result.Error != nil {
		return result.Error
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
func (p *RepoAuthorImpl) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return fmt.Errorf(_emptyRequest)
	}

	var countExist int32
	p.pg.Conn.Model(&entity.Author{}).Where("id = ?", id).Count(&countExist)
	if countExist == 0 {
		return fmt.Errorf("author with given id not exists")
	}

	result := p.pg.Conn.Where("id = ?", id).Delete(&entity.Author{})
	if result.Error != nil {
		return result.Error
	}

	_, err := p.rpkg.Pool.Get().Do("HDEL", (&entity.Author{}).KeyName(), id)
	if err != nil {
		return err
	}

	return nil
}
