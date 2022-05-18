package controllers

import "github.com/gin-gonic/gin"


// GetUsers ... Get all measures
// @Summary Get all measures
// @Description get all measures
// @Tags Measures
// @Success 200 {array} models.Measure
// @Failure 404 {object} object
// @Router /api/measures [get]
func GetMeasures (context *gin.Context) {
   context.JSON(200, gin.H{"msj": "en el index"})
}