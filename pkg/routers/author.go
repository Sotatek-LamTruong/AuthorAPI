package routers

import (
	"book-author/pkg/config"
	"book-author/pkg/handlers"
	"book-author/pkg/repository"
	"book-author/pkg/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthorRoutes(route *gin.RouterGroup) {
	h := handlers.NewAuthorHandlers(services.NewAuthor(repository.NewAuthorRepo(config.DB)))
	route.GET("/list", h.GetAllAuthors())
	fmt.Println("hi")
	route.GET("/author/{id}", h.GetAuthor())

	route.POST("/author/create", h.CreateAuthor())

	route.POST("author/getByBook/{id}", h.GetAuthorByBook())
}
