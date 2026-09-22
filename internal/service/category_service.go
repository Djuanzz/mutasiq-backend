package service

import (
	"errors"
	"strings"

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
	cm.Name = strings.TrimSpace(cm.Name)
	cm.Name = strings.ToLower(cm.Name)

	if cm.Name == "" {
		return errors.New("category name is required")
	}

	existing, err := cs.repo.FindByName(cm.Name)

	if err != nil {
		return err
	}

	if existing != nil {
		return errors.New("category already exists")
	}

	cm.Id = uuid.New()
	return cs.repo.Create(cm)
}

func (cs *CategoryService) GetAllCategories() ([]model.Category, error) {
	return cs.repo.GetAll()
}

func (cs *CategoryService) DeleteAllCategories() error {
	return cs.repo.DeleteAll()
}
