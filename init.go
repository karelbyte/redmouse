package main

import (
	"github.com/joho/godotenv"
	"os"
)

var AppEnv, DbHost, DbUser, DbPassword, DbPort, DbName, AppRouterPort string

func GetEnvironments() {

	err := godotenv.Load()

	if err != nil {
		println("The enviroment file no exis!")
		os.Exit(1)
	}

	AppEnv = os.Getenv("APP_ENV")

	DbHost = os.Getenv("DBHOST")
	DbUser = os.Getenv("DBUSER")
	DbPassword = os.Getenv("DBPASSWORD")
	DbPort = os.Getenv("DBPORT")
	DbName = os.Getenv("DBNAME")

	println("All environments are set!")
}
