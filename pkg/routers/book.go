package routers

import (
	"book-author/pkg/config"
	"book-author/pkg/handlers"
	"book-author/pkg/repository"
	"book-author/pkg/services"

	"github.com/gin-gonic/gin"
)

func BookRoutes(route *gin.RouterGroup) {
	h := handlers.NewBookHandlers(services.NewBook(repository.NewBookRepo(config.DB)))

	route.POST("/create", h.CreateBook())
	route.GET("/getByCate/:id", h.GetBookByCate())
	route.GET("/getByAuthor/:id", h.GetBookByAuthor())
	route.GET("/getByName/:name", h.GetBookByName())
}
