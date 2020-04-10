package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type client struct {
	Db *gorm.DB
}

func New(dbURL string) (*client, error) {
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	return &client{Db: db}, nil

}
func (cl *client) Close() {
	defer cl.Db.Close()
}
