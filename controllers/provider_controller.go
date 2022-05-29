package controllers

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var provider = models.Provider{}

// GetProviders ... Get all providers
// @Summary Get all providers
// @Description Get all providers, acept pagination ?page=1&page_provider=10
// @Tags Providers
// @Success 200 {array} models.Provider
// @Failure 404 {object} object
// @Router /api/providers [get]
func GetProviders(context *gin.Context) {

	var providers []models.Provider
	var count int64

	db.Conect().Model(&provider).Count(&count)
	if result := db.Conect().Model(&provider).Scopes(Paginate(context)).Find(&providers); result.Error != nil {
		context.JSON(404, ErrorToGetData)
		return
	}

	context.JSON(200, gin.H{"data": providers, "total_rows": count})
}

// GetProviders ... Create provider
// @Summary Create provider
// @Description Create provider, ej. {"description": "new provider name"}
// @Tags Providers
// @Success 200 {array} models.Provider
// @Failure 404 {object} object
// @Router /api/providers [post]
func CreateProvider(context *gin.Context) {

	if err := context.BindJSON(&provider); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	if result := db.Conect().Create(&provider); result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}

	context.JSON(200, provider)
}

// GetProviders ... Find provider
// @Summary Find provider
// @Description Find provider
// @Tags Providers
// @Success 200 {array} models.Provider
// @Failure 404 {object} object
// @Router /api/providers/{id} [get]
func GetProviderByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "error": err.Error()})
		return
	}

	providerToFind := models.Provider{ID: id}
	result := db.Conect().Model(&provider).Find(&providerToFind)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}

	if result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}

	context.JSON(200, providerToFind)
}

// GetProviders ... Update provider
// @Summary Update provider
// @Description Update provider ej. {"description": "update provider name"}
// @Tags Providers
// @Success 200 {array} models.Provider
// @Failure 404 {object} object
// @Router /api/providers/{id} [put]
func UpdateProviderByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "error": err.Error()})
		return
	}

	if err := context.BindJSON(&Payload); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	itemToUpdate := models.Provider{ID: id}
	result := db.Conect().Model(&provider).Find(&itemToUpdate);
	if  result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}

	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}


	if err := models.FillModel(&itemToUpdate, Payload); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	db.Conect().Save(&itemToUpdate)

	context.JSON(200, itemToUpdate)

}

// GetProviders ... Delete provider
// @Summary Delete provider
// @Description Delete provider
// @Tags Providers
// @Success 200 {array} models.Provider
// @Failure 404 {object} object
// @Router /api/providers/{id} [delete]
func DeleteProviderByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	itemToDelete := &models.Provider{ID: id}
	result := db.Conect().Model(&provider).Find(&itemToDelete)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}

	if result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}

	db.Conect().Delete(&itemToDelete)

	context.JSON(200, gin.H{"success": "Item delete"})
}
