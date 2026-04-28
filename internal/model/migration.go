package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Transaction{})

	if err != nil {
		return err
	}

	return nil
}
