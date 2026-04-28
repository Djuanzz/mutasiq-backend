package router

import (
	"github.com/Djuanzz/mutasiq-backend/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api")

	HealthRouter(api)
	TransactionRouter(api, db)

	return r
}
