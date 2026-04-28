package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Djuanzz/mutasiq-backend/internal/service"
	"github.com/Djuanzz/mutasiq-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(ts *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: ts,
	}
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
