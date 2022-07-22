package handlers

import (
	"book-author/pkg/dto"
	"book-author/pkg/services"
	"fmt"
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
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		row, err := h.cateServices.CreateCategory(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}
		msg := fmt.Sprintf("new category has id is %d add to db", row)
		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status": "ok",
			"code":   "200",
			"data":   msg,
		})
	}
}

func (h CategoryHandlers) GetCateById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookID := ctx.Param("id")

		id, err := StrToInt(bookID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
			return
		}
		res, err := h.cateServices.GetCate(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}
		// if err == nil && res.CategoryId == 0 {
		// 	fmt.Println("3")
		// 	ctx.JSON(ctx.Writer.Status(), gin.H{
		// 		"status": "Id not exist",
		// 		"code":   400,
		// 	})
		// 	return
		// }
		ctx.Header("content-type", "application/json")
		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status":   "ok",
			"code":     "200",
			"category": res,
		})
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
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}

		ctx.Header("content-type", "application/json")
		ctx.JSON(ctx.Writer.Status(), res)
	}
}

func (h CategoryHandlers) DeleteCate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		input, err := StrToInt(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
				"code":   http.StatusBadRequest,
			})
		}
		row, err := h.cateServices.DeleteCategory(input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
				"code":   http.StatusInternalServerError,
			})
			return
		}

		msg := fmt.Sprintf("a category has id is %d remove to db", row)
		ctx.JSON(ctx.Writer.Status(), gin.H{
			"status": "delete successfully category",
			"code":   "200",
			"data":   msg,
		})
	}
}
