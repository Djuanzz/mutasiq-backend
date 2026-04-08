package router

import (
	"github.com/Djuanzz/mutasiq-backend/internal/handler"
	"github.com/Djuanzz/mutasiq-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func TransactionRouter(r *gin.RouterGroup) {
	service := service.NewTransactionService()
	handler := handler.NewTransactionHandler(service)
	transaction := r.Group("/transaction")

	transaction.POST("/process", handler.ProcessTransactionFile)
	transaction.GET("/blue", handler.BlueTakeJson)
}
