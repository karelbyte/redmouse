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
// @Description get all measures
// @Tags Measures
// @Success 200 {array} models.Measure
// @Failure 404 {object} object
// @Router /api/measures [get]
func GetMeasures(context *gin.Context) {
	var measures []models.Measure
	var count int64
	db.Conect().Model(&measures).Count(&count)
	result := db.Conect().Scopes(Paginate(context)).Find(&measures)
	if result.Error != nil {
		context.JSON(200, gin.H{"error": "A ocurrido un error al obtener los datos."})
		return
	}
	
	context.JSON(200, gin.H{"data": measures, "total_rows": count })
}

// GetUsers ... Create measure
// @Summary Create measure
// @Description create measure
// @Tags Measures
// @Success 200 {array} models.Measure
// @Failure 404 {object} object
// @Router /api/measures [post]
func CreateMeasure(context *gin.Context) {

	if err := context.BindJSON(&measure); err != nil {
		context.JSON(200, gin.H{"error1": err.Error()})
		return
	}
	result := db.Conect().Create(&measure)
	if result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}
	context.JSON(200, measure)
}

// GetUsers ... Find measure
// @Summary Find measure
// @Description find measure
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
// @Description update measure
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
	measureFind := models.Measure{ID: id}
    result := db.Conect().Find(&measureFind)
	if result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}
	if err := context.BindJSON(&measure); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	measureFind.Description = measure.Description
	db.Conect().Save(&measureFind)

	context.JSON(200, measureFind)
}


// GetUsers ... Delete measure
// @Summary Delete measure
// @Description delete measure
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
	measureFind := models.Measure{ID: id}
    result := db.Conect().Find(&measureFind)
	if result.RowsAffected == 0 { 
		context.JSON(200, gin.H{})
		return
	}
	if result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}
	db.Conect().Delete(&measureFind)

	context.JSON(200, gin.H{"success": "Measure delete"})
}
