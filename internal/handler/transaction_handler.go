package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Djuanzz/mutasiq-backend/internal/model"
	"github.com/Djuanzz/mutasiq-backend/internal/service"
	"github.com/Djuanzz/mutasiq-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(ts *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: ts,
	}
}

func (h *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var txn model.Transaction
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateTransaction(&txn); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Transaction created successfully", txn)
}

func (h *TransactionHandler) UpdateTransaction(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid transaction ID")
		return
	}

	var txn model.Transaction
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	txn.Id = id

	if err := h.service.UpdateTransaction(&txn); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Transaction updated successfully", txn)
}

func (h *TransactionHandler) DeleteTransaction(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid transaction ID")
		return
	}

	if err := h.service.DeleteTransaction(id); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Transaction deleted successfully", nil)
}

func (h *TransactionHandler) GetAllTransactions(ctx *gin.Context) {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, "Transactions fetched successfully", transactions)
}

func (h *TransactionHandler) ProcessTransactionFile(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "File is required")
		return
	}

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create directory")
		return
	}

	filePath := filepath.Join(uploadDir, file.Filename)

	// simpan file ke disk
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to save file")
		return
	}

	fmt.Println("File saved at:", filePath)

	// baru kirim path ke service
	transactions, err := h.service.ProcessTransactionFile(filePath, file.Filename)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "File processed successfully", transactions)
}
