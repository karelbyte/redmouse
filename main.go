package main

import (
	"elpuertodigital/redmouse/db"
	"elpuertodigital/redmouse/http"
)


// @title Redmouse API documentation.
// @version 1.0.0
// @BasePath /api/*

func main() {
	GetEnvironments()
	db.ExecMigrations()
	http.Routes()
	println("All ok")
}