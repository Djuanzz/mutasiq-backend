package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id" form:"id"`
	Name string    `gorm:"type:varchar(255)" json:"name" form:"name"`

	Rules []CategoryRule `gorm:"foreignKey:CategoryId" json:"rules" form:"rules"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" form:"updated_at"`
}
