package repository

import (
	"github.com/Djuanzz/mutasiq-backend/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) Create(cm *model.Category) error {
	return cr.db.Create(cm).Error
}
