package controllers

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var client = models.Client{}

// GetClients ... Get all clients
// @Summary Get all clients
// @Description Get all clients, acept pagination ?page=1&page_client=10
// @Tags Clients
// @Success 200 {array} models.Client
// @Failure 404 {object} object
// @Router /api/clients [get]
func GetClients(context *gin.Context) {

	var clients []models.Client
	var count int64

	db.Conect().Model(&client).Count(&count)
	if result := db.Conect().Model(&client).Scopes(Paginate(context)).Find(&clients); result.Error != nil {
		context.JSON(404, ErrorToGetData)
		return
	}

	context.JSON(200, gin.H{"data": clients, "total_rows": count})
}

// GetClients ... Create client
// @Summary Create client
// @Description Create client, ej. {"description": "new client name"}
// @Tags Clients
// @Success 200 {array} models.Client
// @Failure 404 {object} object
// @Router /api/clients [post]
func CreateClient(context *gin.Context) {

	if err := context.BindJSON(&client); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	if result := db.Conect().Create(&client); result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}

	context.JSON(200, client)
}

// GetClients ... Find client
// @Summary Find client
// @Description Find client
// @Tags Clients
// @Success 200 {array} models.Client
// @Failure 404 {object} object
// @Router /api/clients/{id} [get]
func GetClientByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "error": err.Error()})
		return
	}

	clientToFind := models.Client{ID: id}
	result := db.Conect().Model(&client).Find(&clientToFind)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}

	if result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}

	context.JSON(200, clientToFind)
}

// GetClients ... Update client
// @Summary Update client
// @Description Update client ej. {"description": "update client name"}
// @Tags Clients
// @Success 200 {array} models.Client
// @Failure 404 {object} object
// @Router /api/clients/{id} [put]
func UpdateClientByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	if err := context.BindJSON(&Payload); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	itemToUpdate := models.Client{ID: id}

	result := db.Conect().Model(&client).Find(&itemToUpdate)
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

// GetClients ... Delete client
// @Summary Delete client
// @Description Delete client
// @Tags Clients
// @Success 200 {array} models.Client
// @Failure 404 {object} object
// @Router /api/clients/{id} [delete]
func DeleteClientByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "error": err.Error()})
		return
	}

	itemToDelete := &models.Client{ID: id}
	result := db.Conect().Model(&client).Find(&itemToDelete)
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
