package handlers

import (
	"book-author/pkg/dto"
	"book-author/pkg/services"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthorHandlers struct {
	authorServices services.AuthorServices
}

func NewAuthorHandlers(authorServices services.AuthorServices) AuthorHandlers {
	return AuthorHandlers{
		authorServices: authorServices,
	}
}

func (h AuthorHandlers) GetAllAuthors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var res *dto.ListAuthor
		var err error
		res, err = h.authorServices.GetAllAuthors()
		if err != nil {
			panic(err.Error())
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(http.StatusOK, res.Authors)
	}
}

// func (h AuthorHandlers) GetAuthor() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		authID := ctx.Param("id")

// 		var res *dto.GetAuthorRes
// 		var err error
// 		id, erros := StrToInt(authID)
// 		if erros != nil {
// 			log.Fatal(erros)
// 		}
// 		res, err = h.authorServices.GetAuthor(id)

// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		ctx.Header("content-type", "application/json")
// 		ctx.JSON(http.StatusOK, res.Author)
// 	}
// }

func (h AuthorHandlers) CreateAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.CreateAuthorReq
		err := ctx.BindJSON(&req)
		if err != nil {
			log.Fatalln(err)
		}
		erros := h.authorServices.CreateAuthor(req)

		if erros != nil {
			log.Fatalln(erros)
		}
		fmt.Println("Insert success")
	}
}

func (h AuthorHandlers) GetAuthorByBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookID := ctx.Param("id")
		var err error
		id, erros := StrToInt(bookID)
		if erros != nil {
			log.Fatal(erros)
		}
		res, err := h.authorServices.GetAuthorByBook(id)

		if err != nil {
			panic(err.Error())
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(http.StatusOK, res.Author)
	}
}

func StrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
