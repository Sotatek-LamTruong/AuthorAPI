package handlers

import (
	"book-author/pkg/dto"
	"book-author/pkg/errors"
	"book-author/pkg/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
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
		var err *errors.AppError
		res, err = h.authorServices.GetAllAuthors()
		if err != nil {
			panic(err.Error())
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(http.StatusOK, res.Authors)
	}
}

func (h AuthorHandlers) GetAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authID := chi.URLParam(ctx.Request, "id")

		var res *dto.GetAuthorRes
		var err *errors.AppError
		id, erros := StrToInt(authID)
		if erros != nil {
			log.Fatal(erros)
		}
		res, err = h.authorServices.GetAuthor(id)

		out, errors := json.MarshalIndent(res, " ", "     ")

		if errors != nil {
			panic(err.Error())
		}

		fmt.Println(out)
	}
}

func (h AuthorHandlers) CreateAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req *dto.CreateAuthorReq
		var err *errors.AppError

		err = h.authorServices.CreateAuthor(req)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Insert success")
	}
}

func (h AuthorHandlers) GetAuthorByBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *dto.GetAuthorByBookReq
		var err *errors.AppError

		res, err := h.authorServices.GetAuthorByBook(req)

		if err != nil {
			panic(err.Error())
		}

		out, errors := json.MarshalIndent(res, " ", "     ")

		if errors != nil {
			panic(err.Error())
		}
		fmt.Println(out)
	}
}

func StrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
