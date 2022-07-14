package routers

import (
	"book-author/pkg/config"
	"book-author/pkg/handlers"
	"book-author/pkg/repository"
	"book-author/pkg/services"

	"github.com/gin-gonic/gin"
)

func AuthorRoutes(route *gin.RouterGroup) {
	h := handlers.NewAuthorHandlers(services.NewAuthor(repository.NewAuthorRepo(config.DB)))
	route.GET("/list", h.GetAllAuthors())
	route.POST("/create", h.CreateAuthor())
	route.GET("/get/:id", h.GetAuthor())
	route.GET("/getByBook", h.GetAuthorByBook())
}
