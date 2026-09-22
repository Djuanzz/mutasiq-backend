package service

import (
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

func (s *CategoryService) CreateCategory(cm *model.Category) error {
	cm.Id = uuid.New()
	return s.repo.Create(cm)
}
