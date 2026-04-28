package repository

import (
	"github.com/Djuanzz/mutasiq-backend/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transaction *model.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *TransactionRepository) CreateBatch(transactions []model.Transaction) error {
	return r.db.Create(&transactions).Error
}

func (r *TransactionRepository) GetAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) Update(transaction *model.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *TransactionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Transaction{}, id).Error
}
