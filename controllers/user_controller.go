package controllers

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

var user = models.User{}

// GetUsers ... Get all users
// @Summary Get all users
// @Description Get all users, acept pagination ?page=1&page_user=10
// @Tags Users
// @Success 200 {array} models.User
// @Failure 404 {object} object
// @Router /api/users [get]
func GetUsers(context *gin.Context) {

	var users []models.APIUser
	var count int64

	userId := context.GetString("userId")

	fmt.Println(userId)

	db.Conect().Model(&user).Count(&count)
	if result := db.Conect().Model(&user).Scopes(Paginate(context)).Find(&users); result.Error != nil {
		context.JSON(404, ErrorToGetData)
		return
	}

	context.JSON(200, gin.H{"data": users, "total_rows": count})
}

// GetUsers ... Create user
// @Summary Create user
// @Description Create user, ej. {"description": "new user name"}
// @Tags Users
// @Success 200 {array} models.User
// @Failure 404 {object} object
// @Router /api/users [post]
func CreateUser(context *gin.Context) {

	if err := context.BindJSON(&user); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	if result := db.Conect().Create(&user); result.Error != nil {
		context.JSON(200, gin.H{"error": result.Error})
		return
	}
	context.JSON(200, user)
}

// GetUsers ... Find user
// @Summary Find user
// @Description Find user
// @Tags Users
// @Success 200 {array} models.User
// @Failure 404 {object} object
// @Router /api/users/{id} [get]
func GetUserByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	userToFind := models.APIUser{ID: id}
	result := db.Conect().Model(&user).Find(&userToFind)
	if result.RowsAffected == 0 {
		context.JSON(200, gin.H{})
		return
	}
	if result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}

	context.JSON(200, userToFind)
}

// GetUsers ... Update user
// @Summary Update user
// @Description Update user ej. {"description": "update user name"}
// @Tags Users
// @Success 200 {array} models.User
// @Failure 404 {object} object
// @Router /api/users/{id} [put]
func UpdateUserByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	if err := context.BindJSON(&Payload); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	itemToUpdate := models.User{ID: id}
	if result := db.Conect().Model(&user).Find(&itemToUpdate); result.Error != nil {
		context.JSON(200, gin.H{"Error": result.Error})
		return
	}

	if err := models.FillModel(&itemToUpdate, Payload); err != nil {
		context.JSON(200, gin.H{"error": err.Error()})
		return
	}

	db.Conect().Save(&itemToUpdate)

	apiItem := models.APIUser{}
	err = copier.CopyWithOption(&apiItem, &itemToUpdate, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		context.JSON(200, gin.H{"Error": err.Error()})
		return
	}

	context.JSON(200, apiItem)

}

// GetUsers ... Delete user
// @Summary Delete user
// @Description Delete user
// @Tags Users
// @Success 200 {array} models.User
// @Failure 404 {object} object
// @Router /api/users/{id} [delete]
func DeleteUserByID(context *gin.Context) {

	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(200, gin.H{"id": context.Param("id"), "Error": err.Error()})
		return
	}

	itemToDelete := &models.User{ID: id}
	result := db.Conect().Model(&user).Find(&itemToDelete)
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
