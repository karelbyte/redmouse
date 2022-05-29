package controllers

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var size = models.Size{}

// GetUsers ... Get all sizes
// @Summary Get all sizes
// @Description Get all sizes, acept pagination ?page=1&page_size=10
// @Tags Sizes
// @Success 200 {array} models.Size
// @Failure 404 {object} object
// @Router /api/sizes [get]
func GetSizes(context *gin.Context) {
	var sizes []models.Size
	var count int64
	db.Conect().Model(&models.Size{}).Count(&count)

	if result := db.Conect().Scopes(Paginate(context)).Find(&sizes); result.Error != nil {
		context.JSON(404, ErrorToGetData)
		return
	}

	context.JSON(200, gin.H{"data": sizes, "total_rows": count})
}

// GetUsers ... Create size
// @Summary Create size
// @Description Create size, ej. {"description": "new size name"}
// @Tags Sizes
// @Success 200 {array} models.Size
// @Failure 404 {object} object
// @Router /api/sizes [post]
func CreateSize(context *gin.Context) {

	if err := context.BindJSON(&size); err != nil {
		context.JSON(200, gin.H{"error1": err.Error()})
		return
	}
	
	if result := db.Conect().Create(&size); result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}
	context.JSON(200, size)
}

// GetUsers ... Find size
// @Summary Find size
// @Description Find size
// @Tags Sizes
// @Success 200 {array} models.Size
// @Failure 404 {object} object
// @Router /api/sizes/{id} [get]
func GetSizeByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	size = models.Size{ID: id}
	result := db.Conect().Find(&size)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}
	if result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}

	context.JSON(200, size)
}

// GetUsers ... Update size
// @Summary Update size
// @Description Update size ej. {"description": "update size name"}
// @Tags Sizes
// @Success 200 {array} models.Size
// @Failure 404 {object} object
// @Router /api/sizes/{id} [put]
func UpdateSizeByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	itemToUpdate := models.Size{ID: id}
	result := db.Conect().Find(&itemToUpdate);
	if  result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}

	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}
	
	if err := context.BindJSON(&size); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	itemToUpdate.Description = size.Description
	db.Conect().Save(&itemToUpdate)

	context.JSON(200, itemToUpdate)
}

// GetUsers ... Delete size
// @Summary Delete size
// @Description Delete size
// @Tags Sizes
// @Success 200 {array} models.Size
// @Failure 404 {object} object
// @Router /api/sizes/{id} [delete]
func DeleteSizeByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}
	
	itemToDelete := &models.Size{ID: id}
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
