package entity

import (
	"strings"
	"time"
)

type Publisher struct {
	Id        int32     `json:"id,omitempty" validate:"required"`
	Name      string    `json:"name,omitempty" validate:"required"`
	Detail    string    `json:"detail,omitempty" validate:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (p *Publisher) TableName() string {
	return "t_publisher"
}

func (p *Publisher) KeyName() string {
	return strings.Replace(p.TableName(), "_", ":", -1)
}
