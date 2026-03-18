package handler

import (
	"net/http"

	"github.com/Djuanzz/cashlens-backend/internal/service"
	"github.com/Djuanzz/cashlens-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	service *service.HealthService
}

func NewHealthHandler(hs *service.HealthService) *HealthHandler {
	return &HealthHandler{
		service: hs,
	}
}

func (h *HealthHandler) CheckHealth(C *gin.Context) {
	result := h.service.CheckHealth()

	utils.SuccessResponse(C, http.StatusOK, result)
}
