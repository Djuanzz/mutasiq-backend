package service

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/Djuanzz/mutasiq-backend/internal/model"
	"github.com/Djuanzz/mutasiq-backend/internal/parser"
	"github.com/Djuanzz/mutasiq-backend/internal/repository"
	"github.com/google/uuid"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) ProcessTransactionFile(filePath string, fileName string) ([]model.Transaction, error) {

	text, err := parser.ExtractTextFromPDF(filePath)
	if err != nil {
		return nil, err
	}

	year, found := parser.ExtractYearFromText(text)
	if !found {
		log.Println("Year not found! Raw text snippet:")
		if len(text) > 300 {
			log.Println(text[:300])
		} else {
			log.Println(text)
		}
	}
	cleaned := parser.CleanExtractedText(text)
	lines := strings.Split(cleaned, "\n")

	var results []model.Transaction

	for _, line := range lines {
		txn := parser.ParseTransaction(line, year)
		if txn != nil {
			txn.Id = uuid.New()
			results = append(results, *txn)
		}
	}

	if len(results) > 0 {
		if err := s.repo.CreateBatch(results); err != nil {
			return nil, err
		}
	}

	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	outputPath := filepath.Join("transactions", baseName+".json")

	parser.SaveJSON(outputPath, results)

	return results, nil
}

func (s *TransactionService) CreateTransaction(txn *model.Transaction) error {
	txn.Id = uuid.New()
	return s.repo.Create(txn)
}

func (s *TransactionService) UpdateTransaction(txn *model.Transaction) error {
	return s.repo.Update(txn)
}

func (s *TransactionService) DeleteTransaction(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *TransactionService) GetAllTransactions() ([]model.Transaction, error) {
	return s.repo.GetAll()
}
