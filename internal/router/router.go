package router

import (
	"github.com/Djuanzz/mutasiq-backend/internal/middleware"
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
