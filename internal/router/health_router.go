package router

import (
	"github.com/Djuanzz/mutasiq-backend/internal/handler"
	"github.com/Djuanzz/mutasiq-backend/internal/repository"
	"github.com/Djuanzz/mutasiq-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup) {
	repo := repository.NewHealthRepository()
	service := service.NewHealthService(repo)
	handler := handler.NewHealthHandler(service)

	health := r.Group("/health")

	health.GET("/", handler.CheckHealth)
}
