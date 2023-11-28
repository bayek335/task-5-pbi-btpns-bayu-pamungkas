package main

import (
	"log"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/database"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("environment.env"); err != nil {
		log.Fatal(err.Error())
	}
	db := database.DatabaseConnection()
	router.Router(db)
}
