package routers

import "github.com/gin-gonic/gin"

func SetupRoute(app *gin.Engine) {
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "done",
		})
	})
	v1 := app.Group("/author")
	v2 := app.Group("/category")
	CateRoutes(v2)
	AuthorRoutes(v1)
}
