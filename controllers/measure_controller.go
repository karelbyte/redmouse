package controllers

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var measure = models.Measure{}

// GetUsers ... Get all measures
// @Summary Get all measures
// @Description Get all measures, acept pagination ?page=1&page_size=10
// @Tags Measures
// @Success 200 {array} models.Measure
// @Failure 404 {object} object
// @Router /api/measures [get]
func GetMeasures(context *gin.Context) {
	var measures []models.Measure
	var count int64
	db.Conect().Model(&models.Measure{}).Count(&count)

	if result := db.Conect().Scopes(Paginate(context)).Find(&measures); result.Error != nil {
		context.JSON(404, ErrorToGetData)
		return
	}

	context.JSON(200, gin.H{"data": measures, "total_rows": count})
}

// GetUsers ... Create measure
// @Summary Create measure
// @Description Create measure, ej. {"description": "new measure name"}
// @Tags Measures
// @Success 200 {array} models.Measure
// @Failure 404 {object} object
// @Router /api/measures [post]
func CreateMeasure(context *gin.Context) {

	if err := context.BindJSON(&measure); err != nil {
		context.JSON(200, gin.H{"error1": err.Error()})
		return
	}
	
	if result := db.Conect().Create(&measure); result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}
	context.JSON(200, measure)
}

// GetUsers ... Find measure
// @Summary Find measure
// @Description Find measure
// @Tags Measures
// @Success 200 {array} models.Measure
// @Failure 404 {object} object
// @Router /api/measures/{id} [get]
func GetMeasureByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	measure = models.Measure{ID: id}
	result := db.Conect().Find(&measure)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}
	if result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}

	context.JSON(200, measure)
}

// GetUsers ... Update measure
// @Summary Update measure
// @Description Update measure ej. {"description": "update measure name"}
// @Tags Measures
// @Success 200 {array} models.Measure
// @Failure 404 {object} object
// @Router /api/measures/{id} [put]
func UpdateMeasureByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}
	itemToUpdate := models.Measure{ID: id}
	if result := db.Conect().Find(&itemToUpdate); result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}
	if err := context.BindJSON(&measure); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	itemToUpdate.Description = measure.Description
	db.Conect().Save(&itemToUpdate)

	context.JSON(200, itemToUpdate)
}

// GetUsers ... Delete measure
// @Summary Delete measure
// @Description Delete measure
// @Tags Measures
// @Success 200 {array} models.Measure
// @Failure 404 {object} object
// @Router /api/measures/{id} [delete]
func DeleteMeasureByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}
	itemToDelete := &models.Measure{ID: id}
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
