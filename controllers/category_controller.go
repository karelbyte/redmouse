package controllers

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var category = models.Category{}

// GetUsers ... Get all categories
// @Summary Get all categories
// @Description Get all categories, acept pagination ?page=1&page_size=10
// @Tags Categories
// @Success 200 {array} models.Category
// @Failure 404 {object} object
// @Router /api/categories [get]
func GetCategories(context *gin.Context) {
	var categories []models.Category
	var count int64
	db.Conect().Model(&models.Category{}).Count(&count)

	if result := db.Conect().Scopes(Paginate(context)).Find(&categories); result.Error != nil {
		context.JSON(404, ErrorToGetData)
		return
	}

	context.JSON(200, gin.H{"data": categories, "total_rows": count})
}

// GetUsers ... Create category
// @Summary Create category
// @Description Create category, ej. {"description": "new category name"}
// @Tags Categories
// @Success 200 {array} models.Category
// @Failure 404 {object} object
// @Router /api/categories [post]
func CreateCategory(context *gin.Context) {

	if err := context.BindJSON(&category); err != nil {
		context.JSON(200, gin.H{"error1": err.Error()})
		return
	}
	
	if result := db.Conect().Create(&category); result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}
	context.JSON(200, category)
}

// GetUsers ... Find category
// @Summary Find category
// @Description Find category
// @Tags Categories
// @Success 200 {array} models.Category
// @Failure 404 {object} object
// @Router /api/categories/{id} [get]
func GetCategoryByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	category = models.Category{ID: id}
	result := db.Conect().Find(&category)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}
	if result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}

	context.JSON(200, category)
}

// GetUsers ... Update category
// @Summary Update category
// @Description Update category ej. {"description": "update category name"}
// @Tags Categories
// @Success 200 {array} models.Category
// @Failure 404 {object} object
// @Router /api/categories/{id} [put]
func UpdateCategoryByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}
	itemToUpdate := models.Category{ID: id}
	if result := db.Conect().Find(&itemToUpdate); result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}
	if err := context.BindJSON(&category); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}
	itemToUpdate.Description = category.Description
	db.Conect().Save(&itemToUpdate)

	context.JSON(200, itemToUpdate)
}

// GetUsers ... Delete category
// @Summary Delete category
// @Description Delete category
// @Tags Categories
// @Success 200 {array} models.Category
// @Failure 404 {object} object
// @Router /api/categories/{id} [delete]
func DeleteCategoryByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}
	itemToDelete := &models.Category{ID: id}
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
