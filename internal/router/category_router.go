package router

import (
	"github.com/Djuanzz/mutasiq-backend/internal/handler"
	"github.com/Djuanzz/mutasiq-backend/internal/repository"
	"github.com/Djuanzz/mutasiq-backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRouter(r *gin.RouterGroup, db *gorm.DB) {
	cr := repository.NewCategoryRepository(db)
	cs := service.NewCategoryService(cr)
	ch := handler.NewCategoryHandler(cs)
	category := r.Group("/category")

	category.POST("/", ch.CreateCategory)
}
