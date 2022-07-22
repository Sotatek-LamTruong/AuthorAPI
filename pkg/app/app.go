package app

import (
	"book-author/pkg/config"
	"book-author/pkg/routers"

	"github.com/gin-gonic/gin"
)

func SetupApp() *gin.Engine {
	config.Init()
	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	routers.SetupRoute(app)

	return app
}
