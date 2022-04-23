package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	Id        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
