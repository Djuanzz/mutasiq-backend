package service

import (
	"github.com/Djuanzz/mutasiq-backend/internal/model"
	"github.com/Djuanzz/mutasiq-backend/internal/repository"
)

type HealthService struct {
	repo *repository.HealthRepository
}

func NewHealthService(repo *repository.HealthRepository) *HealthService {
	return &HealthService{
		repo: repo}
}

func (s *HealthService) CheckHealth() model.Health {
	return s.repo.GetHealth()
}
