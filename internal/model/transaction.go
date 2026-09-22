package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id" form:"id"`
	Date   string    `gorm:"type:varchar(255)" json:"date" form:"date"`
	Amount float64   `gorm:"type:decimal(10,2)" json:"amount" form:"amount"`
	Type   string    `gorm:"type:varchar(255)" json:"type" form:"type"`
	Desc   string    `gorm:"type:varchar(255)" json:"desc" form:"desc"`

	CategoryId *uuid.UUID `gorm:"type:uuid" json:"category_id" form:"category_id"`
	Category   *Category  `gorm:"foreignKey:CategoryId" json:"category" form:"category"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" form:"updated_at"`
}
