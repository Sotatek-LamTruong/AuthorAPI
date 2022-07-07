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
	route.GET("/", h.GetAllAuthors())

	route.GET("/author/{id}", h.GetAuthor())

	route.POST("/author/create", h.CreateAuthor())

	route.POST("author/getByBook/{id}", h.GetAuthorByBook())
}
