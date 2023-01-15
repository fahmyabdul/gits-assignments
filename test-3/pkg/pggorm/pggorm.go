package pggorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PgGorm :
type PgGorm struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DB     string `json:"db"`
	Schema string `json:"schema"`
	Conn   *gorm.DB
}

// NewPostgreGORM :
func NewPostgreGORM(host string, port int, user, pass, db, schema string) (*PgGorm, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable", host, port, user, pass, db, schema)

	gormConnPostgre, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	gormConf := PgGorm{
		Host:   host,
		Port:   port,
		User:   user,
		DB:     db,
		Schema: schema,
		Conn:   gormConnPostgre,
	}

	return &gormConf, nil
}

// Close :
func (a *PgGorm) Close() error {
	err := a.Conn.Close()

	return err
}
