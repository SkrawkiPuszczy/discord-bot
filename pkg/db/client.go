package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type client struct {
	db *gorm.DB
}

func New(dbURL string) (*client, error) {
	db, err := gorm.Open("postgres", dbURL)
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Adv{})
	return &client{db: db}, nil

}
func (cl *client) Close() {
	defer cl.db.Close()
}
