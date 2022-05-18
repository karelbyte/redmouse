package http

import (
	"os"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" 
    "github.com/swaggo/files"
	_"elpuertodigital/redmouse/docs"
)

var api *gin.RouterGroup

func Routes () {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.LoadHTMLGlob("./public/*.html")

	router.Use(static.Serve("/assets", static.LocalFile("./public/assets", false)))

	api = router.Group("/api")

	ManagerApiRoutes(api);

	router.GET("/help/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// This code is use to catch all routes that not match with "/apis" and respose front app
	router.NoRoute(ManagerFrontRoutes)

	appRouterPort := os.Getenv("APP_ROUTER_PORT")
	router.Run(":" + appRouterPort)
}