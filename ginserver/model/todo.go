package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Id     uuid.UUID `gorm:"type:uuid"`
	Title  string    `json:"title"`
	Task   string    `json:"task"`
	Finish bool      `json:"finish"`
}
