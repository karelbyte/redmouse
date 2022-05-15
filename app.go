package main

import "elpuertodigital/redmouse/db"

func main() {
	GetEnvironments()
	db.ExecMigrations(AppEnv)
	println("all ok")
}