package models

import (
	"time"
)

type ToDo struct {
	ID        uint64     `gorm:"primaryKey"`
	CreatedAt *time.Time `gorm:"type:datetime(3)"`
	UpdatedAt *time.Time `gorm:"type:datetime(3)"`
	DeletedAt *time.Time `gorm:"index;type:datetime(3)"`
	Title     string     `gorm:"type:longtext"`
	Done      bool       `gorm:"type:tinyint(1)"`
}

func (t *ToDo) TableName() string {
	return "todos"
}
