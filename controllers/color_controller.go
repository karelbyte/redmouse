package controllers

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var color = models.Color{}

// GetUsers ... Get all colors
// @Summary Get all colors
// @Description Get all colors, acept pagination ?page=1&page_color=10
// @Tags Colors
// @Success 200 {array} models.Color
// @Failure 404 {object} object
// @Router /api/colors [get]
func GetColors(context *gin.Context) {
	var colors []models.Color
	var count int64
	db.Conect().Model(&models.Color{}).Count(&count)

	if result := db.Conect().Scopes(Paginate(context)).Find(&colors); result.Error != nil {
		context.JSON(404, ErrorToGetData)
		return
	}

	context.JSON(200, gin.H{"data": colors, "total_rows": count})
}

// GetUsers ... Create color
// @Summary Create color
// @Description Create color, ej. {"description": "new color name"}
// @Tags Colors
// @Success 200 {array} models.Color
// @Failure 404 {object} object
// @Router /api/colors [post]
func CreateColor(context *gin.Context) {

	if err := context.BindJSON(&color); err != nil {
		context.JSON(200, gin.H{"error1": err.Error()})
		return
	}
	
	if result := db.Conect().Create(&color); result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}

	context.JSON(200, color)
}

// GetUsers ... Find color
// @Summary Find color
// @Description Find color
// @Tags Colors
// @Success 200 {array} models.Color
// @Failure 404 {object} object
// @Router /api/colors/{id} [get]
func GetColorByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	color = models.Color{ID: id}
	result := db.Conect().Find(&color)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}

	if result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}

	context.JSON(200, color)
}

// GetUsers ... Update color
// @Summary Update color
// @Description Update color ej. {"description": "update color name"}
// @Tags Colors
// @Success 200 {array} models.Color
// @Failure 404 {object} object
// @Router /api/colors/{id} [put]
func UpdateColorByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	itemToUpdate := models.Color{ID: id}
	result := db.Conect().Find(&itemToUpdate);
	if  result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}

	if err := context.BindJSON(&color); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	itemToUpdate.Description = color.Description
	db.Conect().Save(&itemToUpdate)

	context.JSON(200, itemToUpdate)
}

// GetUsers ... Delete color
// @Summary Delete color
// @Description Delete color
// @Tags Colors
// @Success 200 {array} models.Color
// @Failure 404 {object} object
// @Router /api/colors/{id} [delete]
func DeleteColorByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	itemToDelete := &models.Color{ID: id}
	result := db.Conect().Find(&itemToDelete)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}

	if result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}
	
	db.Conect().Delete(&itemToDelete)

	context.JSON(200, gin.H{"success": "Item delete"})
}
