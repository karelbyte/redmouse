package controllers

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

var product = models.Product{}

// GetProducts ... Get all products
// @Summary Get all products
// @Description Get all products, acept pagination ?page=1&page_product=10
// @Tags Products
// @Success 200 {array} models.Product
// @Failure 404 {object} object
// @Router /api/products [get]
func GetProducts(context *gin.Context) {

	var products []models.Product
	var count int64

	db.Conect().Model(&product).Count(&count)
	result := db.Conect().Preload(clause.Associations).Model(&product).Scopes(Paginate(context)).Find(&products)
	if  result.Error != nil {
		context.JSON(404, ErrorToGetData)
		return
	}

	context.JSON(200, gin.H{"data": products, "total_rows": count})
}

// GetProducts ... Create product
// @Summary Create product
// @Description Create product, ej. {"description": "new product name"}
// @Tags Products
// @Success 200 {array} models.Product
// @Failure 404 {object} object
// @Router /api/products [post]
func CreateProduct(context *gin.Context) {

	if err := context.BindJSON(&product); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	// if result := db.Conect().Model(&models.Product{}).Omit(clause.Associations).Create(&product); result.Error != nil {
	// 	context.JSON(200, gin.H{"error": result.Error})
	// 	return
	// }

	if result := db.Conect().Model(&models.Product{}).Omit("Measure").Omit("Provider").Omit("Category").Create(&product); result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}

	context.JSON(200, product)
}

// GetProducts ... Find product
// @Summary Find product
// @Description Find product
// @Tags Products
// @Success 200 {array} models.Product
// @Failure 404 {object} object
// @Router /api/products/{id} [get]
func GetProductByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "error": err.Error()})
		return
	}

	productToFind := models.Product{ID: id}
	result := db.Conect().Preload(clause.Associations).Model(&product).Find(&productToFind)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}

	if result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}

	context.JSON(200, productToFind)
}

// GetProducts ... Update product
// @Summary Update product
// @Description Update product ej. {"description": "update product name"}
// @Tags Products
// @Success 200 {array} models.Product
// @Failure 404 {object} object
// @Router /api/products/{id} [put]
func UpdateProductByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	if err := context.BindJSON(&Payload); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	itemToUpdate := models.Product{ID: id}

	result := db.Conect().Preload(clause.Associations).Model(&product).Find(&itemToUpdate)
	if result.Error != nil {
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

// GetProducts ... Delete product
// @Summary Delete product
// @Description Delete product
// @Tags Products
// @Success 200 {array} models.Product
// @Failure 404 {object} object
// @Router /api/products/{id} [delete]
func DeleteProductByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "error": err.Error()})
		return
	}

	itemToDelete := &models.Product{ID: id}
	result := db.Conect().Model(&product).Find(&itemToDelete)
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
