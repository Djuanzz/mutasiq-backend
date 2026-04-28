package model

import "github.com/google/uuid"

type Transaction struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id" form:"id"`
	Date     string    `gorm:"type:varchar(255)" json:"date" form:"date"`
	Amount   float64   `gorm:"type:decimal(10,2)" json:"amount" form:"amount"`
	Type     string    `gorm:"type:varchar(255)" json:"type" form:"type"`
	Desc     string    `gorm:"type:varchar(255)" json:"desc" form:"desc"`
	Category *string   `gorm:"type:varchar(255)" json:"category" form:"category"`
}
