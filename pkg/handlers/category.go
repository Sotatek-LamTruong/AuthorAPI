package handlers

import (
	"book-author/pkg/dto"
	"book-author/pkg/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type CategoryHandlers struct {
	cateServices services.CategoryServices
}

func NewCategoryHandlers(cateServices services.CategoryServices) CategoryHandlers {
	return CategoryHandlers{
		cateServices: cateServices,
	}
}

func (h CategoryHandlers) CreateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.CreateCateReq
		var err error
		errs := ctx.BindJSON(&req)
		if errs != nil {
			fmt.Println("Convert fail")
			return
		}
		err = h.cateServices.CreateCategory(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.Writer.Status()
		log.Println("Insert success")
	}
}

func (h CategoryHandlers) GetCateById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookID := ctx.Param("id")

		id, err := StrToInt(bookID)
		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := h.cateServices.GetCateById(id)

		if err != nil {
			fmt.Println(err)
			return
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(ctx.Writer.Status(), res)
	}
}

func (h CategoryHandlers) GetCateByBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var err error
		fmt.Println(id)
		input, err := StrToInt(id)
		if err != nil {
			fmt.Println("Error")
			return
		}
		res, err := h.cateServices.GetCateByBook(input)

		if err != nil {
			fmt.Println("Fail")
			return
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(ctx.Writer.Status(), res)
	}
}

func (h CategoryHandlers) DeleteCate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var err error
		input, err := StrToInt(id)
		if err != nil {
			fmt.Println("Error")
			return
		}
		err = h.cateServices.DeleteCategory(input)
		if err != nil {
			fmt.Println("Fail")
			return
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(ctx.Writer.Status(), "Delete success")
	}
}
