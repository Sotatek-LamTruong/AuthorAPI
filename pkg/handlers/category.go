package handlers

import (
	"book-author/pkg/dto"
	"book-author/pkg/errors"
	"book-author/pkg/services"
	"fmt"
	"log"
	"net/http"

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
		}
		err = h.cateServices.CreateCategory(&req)
		if err != nil {
			panic(err.Error())
		}
		ctx.Writer.Status()
		log.Println("Insert success")
	}
}

func (h CategoryHandlers) GetCateById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookID := ctx.Param("id")

		var errResp *errors.AppError
		id, err := StrToInt(bookID)
		if err != nil {
			errResp.Message = err.Error()
		}
		res, err := h.cateServices.GetCateById(id)

		if err != nil {
			errResp.Message = err.Error()
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(http.StatusOK, res)
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
		}
		res, err := h.cateServices.GetCateByBook(input)

		if err != nil {
			fmt.Println("Fail")
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(http.StatusOK, res)
	}
}
