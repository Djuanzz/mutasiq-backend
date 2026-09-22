package repository

import (
	"errors"

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

func (cr *CategoryRepository) FindByName(name string) (*model.Category, error) {
	var category model.Category

	err := cr.db.Where("name = ?", name).First(&category).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (cr *CategoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	err := cr.db.Find(&categories).Error

	return categories, err
}
