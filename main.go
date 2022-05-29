package main

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/http"

	"github.com/gin-gonic/gin"
)

// @title Redmouse API documentation.
// @version 1.0.0
// @BasePath /api/*

func RespondWithError(contex *gin.Context, code int, message interface{}) {
	contex.AbortWithStatusJSON(code, gin.H{"error": message})
}

func main() {
	GetEnvironments()
	db.ExecMigrations()
	http.Routes()
	println("All ok")
}