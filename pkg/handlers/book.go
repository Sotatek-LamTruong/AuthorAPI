package handlers

import (
	"book-author/pkg/dto"
	"book-author/pkg/services"
	"fmt"

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
		var err error
		errs := ctx.BindJSON(&req)
		if errs != nil {
			fmt.Println("Convert fail")
			return
		}
		err = b.bookServices.CreateBook(&req)
		if err != nil {
			fmt.Println("Create Fail")
			return
		}

		ctx.JSON(ctx.Writer.Status(), "Create success")
	}
}

func (b BookHandlers) GetBookByAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authId := ctx.Param("id")

		// var errResp error
		id, err := StrToInt(authId)
		if err != nil {
			fmt.Println("Convert fail")
			return
		}
		res, err := b.bookServices.GetBookByAuthor(id)
		if err != nil {
			fmt.Println("Fail")
			return
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(ctx.Writer.Status(), res.Books)
	}
}

func (b BookHandlers) GetBookByCate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cateId := ctx.Param("id")

		// var errResp error
		id, err := StrToInt(cateId)
		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := b.bookServices.GetBookByCate(id)
		if err != nil {
			fmt.Println("Fail")
			return
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(ctx.Writer.Status(), res.Books)
	}
}

func (b BookHandlers) GetBookByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nameAuthor := ctx.Param("name")

		res, err := b.bookServices.GetBookByName(nameAuthor)
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(ctx.Writer.Status(), res.Book)
	}
}

func (b BookHandlers) UpdateAuthorByBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.UpdateAuthorByBookReq
		aId, _ := StrToInt(ctx.Param("authid"))
		bId, _ := StrToInt(ctx.Param("bookid"))
		err := ctx.BindJSON(&req)
		if err != nil {
			fmt.Println("Get data fail")
			return
		}
		b.bookServices.UpdateAuthorByBook(aId, bId, &req)
		ctx.Header("content-type", "application/json")
		ctx.JSON(ctx.Writer.Status(), "Update Success")
	}
}
