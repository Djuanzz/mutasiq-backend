package repository

import "github.com/Djuanzz/mutasiq-backend/internal/model"

type HealthRepository struct {
}

// --- constructor function untuk membuat instance baru dari HealthRepository
// --- retrun pointer karena nanti akan di assign ke struct handler
// ---biar ga terjadi copy paste struct handler nya
func NewHealthRepository() *HealthRepository {
	// --- return & karena return type nya pointer
	return &HealthRepository{}
}

// --- method receiver artinya method milik struct HealthRepository
func (r *HealthRepository) GetHealth() model.Health {
	return model.Health{
		Status: "OK",
	}
}
