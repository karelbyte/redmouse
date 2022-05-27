package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Model interface {
	Fill(data interface{})
}

var Payload interface{}

var ErrorToGetData = gin.H{"error": "A ocurrido un error al obtener los datos."}

func Paginate(context *gin.Context) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))

		println(page)
		println(pageSize)

		if int(page) == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
