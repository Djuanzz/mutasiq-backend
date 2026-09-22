package model

import (
	"time"

	"github.com/google/uuid"
)

type CategoryRule struct {
	Id         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id" form:"id"`
	CategoryId uuid.UUID `gorm:"type:uuid" json:"category_id" form:"category_id"`
	Rule       string    `gorm:"type:varchar(255)" json:"rule" form:"rule"`
	Keywoyrd   string    `gorm:"type:varchar(255)" json:"keyword" form:"keyword"`
	Priority   int       `gorm:"type:int" json:"priority" form:"priority"`

	Category Category `gorm:"foreignKey:CategoryId" json:"category" form:"category"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" form:"updated_at"`
}
