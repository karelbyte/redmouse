package http

import (
	"elpuertodigital/redmouse/controllers"

	"github.com/gin-gonic/gin"
)

func ManagerApiRoutes(api *gin.RouterGroup) {

	measure := api.Group("/measures")
	{
		measure.GET("/", controllers.GetMeasures)
		// measure.POST("/", controller.CreateUser)
		// measure.GET("/:id", controller.GetUserByID)
	}

}
