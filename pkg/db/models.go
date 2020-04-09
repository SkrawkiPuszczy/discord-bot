package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Base struct {
	ID        string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4().String()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

type Adv struct {
	gorm.Model
	Name string
}
