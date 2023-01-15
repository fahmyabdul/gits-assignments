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

type RepoPublisherImpl struct {
	pg   *pggorm.PgGorm
	rpkg *redispkg.RedisPkg
}

// NewRepoPublisherImpl : Initialize new Publisher Repository
func NewRepoPublisherImpl(pg *pggorm.PgGorm, rpkg *redispkg.RedisPkg) *RepoPublisherImpl {
	return &RepoPublisherImpl{
		pg:   pg,
		rpkg: rpkg,
	}
}

// Create :
func (p *RepoPublisherImpl) Create(ctx context.Context, request *entity.Publisher) error {
	if request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	// var countExist int32
	// p.pg.Conn.Model(&entity.Publisher{}).Where("name = ?", request.Name).Count(&countExist)
	// if countExist > 0 {
	// 	return fmt.Errorf("publisher with same name already exists")
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
func (p *RepoPublisherImpl) FetchById(ctx context.Context, id int) (*entity.Publisher, error) {
	if id == 0 {
		return nil, fmt.Errorf(_emptyRequest)
	}

	var output entity.Publisher

	result := p.pg.Conn.Where("id = ?", id).Find(&output)
	if result.Error != nil {
		return nil, result.Error
	}

	return &output, nil
}

// FetchByName :
func (p *RepoPublisherImpl) FetchByName(ctx context.Context, name string) ([]*entity.Publisher, error) {
	if name == "" {
		return nil, fmt.Errorf(_emptyRequest)
	}

	var output []*entity.Publisher

	result := p.pg.Conn.Where("name LIKE ?", "%"+name+"%").Find(&output)
	if result.Error != nil {
		return nil, result.Error
	}

	return output, nil
}

// FetchAll :
func (p *RepoPublisherImpl) FetchAll(ctx context.Context) ([]*entity.Publisher, error) {
	var outputs []*entity.Publisher

	result := p.pg.Conn.Find(&outputs)
	if result.Error != nil {
		return nil, result.Error
	}

	return outputs, nil
}

// Update :
func (p *RepoPublisherImpl) Update(ctx context.Context, id int, request *entity.Publisher) error {
	if id == 0 || request == nil {
		return fmt.Errorf(_emptyRequest)
	}

	var countExist int32
	p.pg.Conn.Model(&entity.Publisher{}).Where("id = ?", id).Count(&countExist)
	if countExist == 0 {
		return fmt.Errorf("publisher with given id not exists")
	}

	request.CreatedAt = time.Time{}
	request.UpdatedAt = time.Now()

	result := p.pg.Conn.Model(&entity.Publisher{}).Where("id = ?", id).Updates(request)
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
func (p *RepoPublisherImpl) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return fmt.Errorf(_emptyRequest)
	}

	var countExist int32
	p.pg.Conn.Model(&entity.Publisher{}).Where("id = ?", id).Count(&countExist)
	if countExist == 0 {
		return fmt.Errorf("publisher with given id not exists")
	}

	result := p.pg.Conn.Where("id = ?", id).Delete(&entity.Publisher{})
	if result.Error != nil {
		return result.Error
	}

	_, err := p.rpkg.Pool.Get().Do("HDEL", (&entity.Publisher{}).KeyName(), id)
	if err != nil {
		return err
	}

	return nil
}
