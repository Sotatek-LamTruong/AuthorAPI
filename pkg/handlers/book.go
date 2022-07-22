package handlers

import (
	"book-author/pkg/dto"
	"book-author/pkg/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandlers struct {
	bookServices services.BookServices
}

func NewBookHandlers(bookServices services.BookServices) BookHandlers {
	return BookHandlers{
		bookServices: bookServices,
	}
}

func (b BookHandlers) CreateBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.AddBookReq
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		row, err := b.bookServices.CreateBook(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}

		msg := fmt.Sprintf("new book has id is %d add to db", row)
		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status": "ok",
			"code":   "200",
			"data":   msg,
		})
	}
}

func (b BookHandlers) AddCate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.CateReq
		bookId := ctx.Param("id")
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		id, err := StrToInt(bookId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		res, err := b.bookServices.AddCategory(id, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}

		msg := fmt.Sprintf("new author added to book has id is %d", res.BookID)
		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status": "ok",
			"code":   "200",
			"data":   msg,
		})
	}
}

func (b BookHandlers) AddAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.AuthorReq
		bookId := ctx.Param("id")
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		id, err := StrToInt(bookId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		res, err := b.bookServices.AddAuthors(id, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}

		msg := fmt.Sprintf("new author added to book has id is %d", res.BookID)
		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status": "ok",
			"code":   "200",
			"data":   msg,
		})
	}
}

func (b BookHandlers) EditAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.AuthorReq
		cateId := ctx.Param("id")
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		id, err := StrToInt(cateId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		res, err := b.bookServices.EditAuthor(id, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}

		msg := fmt.Sprintf("%d row updated", res.Row)
		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status": "ok",
			"code":   "200",
			"data":   msg,
		})
	}
}

func (b BookHandlers) DeleteAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.AuthorReq
		authId := ctx.Param("id")
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		id, err := StrToInt(authId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		res, err := b.bookServices.DeleteAuthor(id, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}

		msg := fmt.Sprintf("%d row deleted", res.Row)
		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status": "ok",
			"code":   "200",
			"data":   msg,
		})
	}
}

func (b BookHandlers) GetBooksByAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.AuthorReq
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		res, err := b.bookServices.GetBooksByAuthor(&req)
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

func (b BookHandlers) GetBooksByCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.CateReq
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		res, err := b.bookServices.GetBooksByCate(&req)
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
