package entity

import (
	"strings"
	"time"
)

type Author struct {
	Id        int32     `json:"id,omitempty" validate:"required"`
	Name      string    `json:"name,omitempty" validate:"required"`
	Detail    string    `json:"detail,omitempty" validate:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (p *Author) TableName() string {
	return "t_author"
}

func (p *Author) KeyName() string {
	return strings.Replace(p.TableName(), "_", ":", -1)
}
