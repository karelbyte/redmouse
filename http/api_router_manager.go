package http

import (
	"elpuertodigital/redmouse/controllers"

	"github.com/gin-gonic/gin"
)

func ManagerApiRoutes(api *gin.RouterGroup) {

	measure := api.Group("/measures")
	{
		measure.GET("/", controllers.GetMeasures)
		measure.POST("/", controllers.CreateMeasure)
		measure.GET("/:id", controllers.GetMeasureByID)
		measure.PUT("/:id", controllers.UpdateMeasureByID)
		measure.DELETE("/:id", controllers.DeleteMeasureByID)
	}

}
