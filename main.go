package main

import (
	"github.com/joho/godotenv"
	"github.com/saurabhkanawade/book_gorm_golang_project/database"
	"github.com/saurabhkanawade/book_gorm_golang_project/helper"
	"github.com/saurabhkanawade/book_gorm_golang_project/router"
	"log"
)

func main() {

	log.Println("Starting the application..............")
	err := godotenv.Load("app.env")
	helper.HandleNilError(err)

	//connect to the db
	database.Connection()

	router.Router()
}
