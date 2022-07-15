package handlers

import (
	"book-author/pkg/dto"
	"book-author/pkg/services"
	"fmt"
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
			ctx.JSON(http.StatusExpectationFailed, "Fail get author")
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(http.StatusOK, res.Authors)
	}
}

func (h AuthorHandlers) GetAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authID := ctx.Param("id")

		var res *dto.GetAuthorRes
		var err error
		id, erros := StrToInt(authID)
		if erros != nil {
			fmt.Println("Convert fail")
		}
		res, err = h.authorServices.GetAuthor(id)

		if err != nil {
			ctx.JSON(http.StatusExpectationFailed, "Fail get author")
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(http.StatusOK, res)
	}
}

func (h AuthorHandlers) CreateAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.CreateAuthorReq
		err := ctx.BindJSON(&req)
		if err != nil {
			fmt.Println("Request fail")
		}
		erros := h.authorServices.CreateAuthor(req)

		if erros != nil {
			ctx.JSON(http.StatusExpectationFailed, "Create fail")
		}
		fmt.Println("Insert success")
	}
}

func (h AuthorHandlers) GetAuthorByBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var res *dto.GetAuthorsByBook
		var err error
		res, err = h.authorServices.GetAuthorsByBook()
		if err != nil {
			ctx.JSON(http.StatusExpectationFailed, "Get fail")
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(http.StatusOK, res.Authors)
	}
}

func StrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
