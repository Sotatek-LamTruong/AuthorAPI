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
	route.POST("/category/create/:id", h.AddCate())
	route.POST("/author/create/:id", h.AddAuthor())
	route.PUT("/author/edit/:id", h.EditAuthor())
	route.DELETE("/author/delete/:id", h.DeleteAuthor())
	route.GET("/getByAuthor", h.GetBooksByAuthor())
	route.GET("/getByCate", h.GetBooksByCategory())
}
