package models

import (
	"gorm.io/gorm"
	"time"
)

type Tabler interface {
	TableName() string
}

type Todo struct {
	ID    				uint       		`gorm:"primaryKey"`
	Title				string			`gorm:"column:title"`
	CompletedStatus 	bool			`gorm:"column:completed_status"`
	CreatedAt			time.Time		`gorm:"column:created_at"`
	UpdatedAt			time.Time		`gorm:"column:updated_at"`
	DeletedAt			gorm.DeletedAt	`gorm:"column:deleted_at"`
}

func (Todo) TableName() string {
	return "pm_todos"  // needs to take from config
}

