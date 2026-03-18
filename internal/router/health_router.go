package router

import (
	"github.com/Djuanzz/cashlens-backend/internal/handler"
	"github.com/Djuanzz/cashlens-backend/internal/repository"
	"github.com/Djuanzz/cashlens-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup) {
	repo := repository.NewHealthRepository()
	service := service.NewHealthService(repo)
	handler := handler.NewHealthHandler(service)

	health := r.Group("/health")

	health.GET("/", handler.CheckHealth)
}
