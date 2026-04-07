package router

import (
	"github.com/Djuanzz/cashlens-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api")

	HealthRouter(api)
	TransactionRouter(api)

	return r
}
