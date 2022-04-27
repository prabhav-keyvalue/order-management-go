package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	Id        string         `gorm:"primaryKey" json:"id" example:"951c54e9-4b64-42fe-9d56-e8a9babc3f89"`
	CreatedAt time.Time      `json:"createAt" example:"2022-04-20T17:11:10+05:30"`
	UpdatedAt time.Time      `json:"updatedAt" example:"2022-04-20T17:11:10+05:30"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" swaggertype:"primitive,string" example:"2022-04-20T17:11:10+05:30"`
}
