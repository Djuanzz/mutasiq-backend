package handler

import (
	"net/http"

	"github.com/Djuanzz/mutasiq-backend/internal/model"
	"github.com/Djuanzz/mutasiq-backend/internal/service"
	"github.com/Djuanzz/mutasiq-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(cs *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: cs}
}

func (ch *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var cm model.Category

	if err := ctx.ShouldBind(&cm); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := ch.service.CreateCategory(&cm); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "category created successfully", cm)
}

func (ch *CategoryHandler) GetAllCategories(ctx *gin.Context) {
	categories, err := ch.service.GetAllCategories()

	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "categories fetched successfully", categories)
}

func (ch *CategoryHandler) DeleteAllCategories(ctx *gin.Context) {
	if err := ch.service.DeleteAllCategories(); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "all categories deleted successfully", nil)
}
