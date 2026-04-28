package router

import (
	"github.com/Djuanzz/mutasiq-backend/internal/handler"
	"github.com/Djuanzz/mutasiq-backend/internal/repository"
	"github.com/Djuanzz/mutasiq-backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionRouter(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewTransactionRepository(db)
	svc := service.NewTransactionService(repo)
	handler := handler.NewTransactionHandler(svc)
	transaction := r.Group("/transaction")

	transaction.POST("/upload", handler.ProcessTransactionFile)
	transaction.POST("/", handler.CreateTransaction)
	transaction.GET("/", handler.GetAllTransactions)
	transaction.PATCH("/:id", handler.UpdateTransaction)
	transaction.DELETE("/:id", handler.DeleteTransaction)
}
