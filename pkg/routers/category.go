package routers

import (
	"book-author/pkg/config"
	"book-author/pkg/handlers"
	"book-author/pkg/repository"
	"book-author/pkg/services"

	"github.com/gin-gonic/gin"
)

func CateRoutes(route *gin.RouterGroup) {
	h := handlers.NewCategoryHandlers(services.NewCategory(repository.NewCategoryRepo(config.DB)))

	route.POST("/create", h.CreateCategory())
	route.GET("/getByBook/:id", h.GetCateByBook())
	route.GET("/getByName/:name", h.GetCateByName())

}
