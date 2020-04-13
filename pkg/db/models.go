package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	ID        string     `sql:"primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid.String())
}
