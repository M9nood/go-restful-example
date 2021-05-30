package main

import (
	"log"
	"m9hub/database"
	router "m9hub/routes"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbConfigParams := database.DBConfig{
		DB_USERNAME: os.Getenv("MYSQL_USERNAME"),
		DB_PASSWORD: os.Getenv("MYSQL_PASSWORD"),
		DB_NAME:     os.Getenv("MYSQL_DB_NAME"),
		DB_HOST:     os.Getenv("MYSQL_HOST"),
		DB_PORT:     os.Getenv("MYSQL_PORT"),
	}
	database.InitDb(dbConfigParams)
	router.RouterSetup()
}
