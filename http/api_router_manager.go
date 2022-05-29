package http

import (
	"elpuertodigital/redmouse/controllers"
	"github.com/gin-gonic/gin"
)

func ManagerApiRoutes(api *gin.RouterGroup) {

	auth := api.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
	}

	user := api.Group("/users").Use(controllers.AuthMiddleware())
	{
		user.GET("/", controllers.GetUsers)
		user.POST("/", controllers.CreateUser)
		user.GET("/:id", controllers.GetUserByID)
		user.PUT("/:id", controllers.UpdateUserByID)
		user.DELETE("/:id", controllers.DeleteUserByID)
	}

	measure := api.Group("/measures").Use(controllers.AuthMiddleware())
	{
		measure.GET("/", controllers.GetMeasures)
		measure.POST("/", controllers.CreateMeasure)
		measure.GET("/:id", controllers.GetMeasureByID)
		measure.PUT("/:id", controllers.UpdateMeasureByID)
		measure.DELETE("/:id", controllers.DeleteMeasureByID)
	}

	category := api.Group("/categories").Use(controllers.AuthMiddleware())
	{
		category.GET("/", controllers.GetCategories)
		category.POST("/", controllers.CreateCategory)
		category.GET("/:id", controllers.GetCategoryByID)
		category.PUT("/:id", controllers.UpdateCategoryByID)
		category.DELETE("/:id", controllers.DeleteCategoryByID)
	}

	size := api.Group("/sizes").Use(controllers.AuthMiddleware())
	{
		size.GET("/", controllers.GetSizes)
		size.POST("/", controllers.CreateSize)
		size.GET("/:id", controllers.GetSizeByID)
		size.PUT("/:id", controllers.UpdateSizeByID)
		size.DELETE("/:id", controllers.DeleteSizeByID)
	}

}
