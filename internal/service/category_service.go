package service

import (
	"errors"

	"github.com/Djuanzz/mutasiq-backend/internal/model"
	"github.com/Djuanzz/mutasiq-backend/internal/repository"
	"github.com/google/uuid"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(cr *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: cr}
}

func (cs *CategoryService) CreateCategory(cm *model.Category) error {
	existing, err := cs.repo.FindByName(cm.Name)

	if err != nil {
		return err
	}

	if existing != nil {
		return errors.New("Category already exists")
	}

	cm.Id = uuid.New()
	return cs.repo.Create(cm)
}

func (cs *CategoryService) GetAllCategories() ([]model.Category, error) {
	return cs.repo.GetAll()
}
