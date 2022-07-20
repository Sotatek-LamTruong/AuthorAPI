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
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}

		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status": "ok",
			"code":   "200",
			"list":   res,
		})

	}
}

func (h AuthorHandlers) GetAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authID := ctx.Param("id")
		id, err := StrToInt(authID)

		var res *dto.GetAuthorRes
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		res, err = h.authorServices.GetAuthor(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}

		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status":   "ok",
			"code":     "200",
			"category": res,
		})
	}
}

func (h AuthorHandlers) CreateAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.CreateAuthorReq
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		row, err := h.authorServices.CreateAuthor(req)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}
		msg := fmt.Sprintf("new author has id is %d add to db", row)
		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status": "ok",
			"code":   "200",
			"data":   msg,
		})
	}
}

func (h AuthorHandlers) GetAuthorByBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var res *dto.GetAuthorsByBook
		authID := ctx.Param("id")
		id, err := StrToInt(authID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		res, err = h.authorServices.GetAuthorsByBook(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}
		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status": "ok",
			"code":   "200",
			"res":    res,
		})
	}
}

func StrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
