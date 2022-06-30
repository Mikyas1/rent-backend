package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	UUID      uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey;column:id"`
	CreatedAt time.Time      `json:"created_at" faker:"-"`
	UpdatedAt time.Time      `json:"-" faker:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" faker:"-"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
