package database

import (
	"github.com/saurabhkanawade/book_gorm_golang_project/helper"
	"github.com/saurabhkanawade/book_gorm_golang_project/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func Connection() {

	log.Println("connecting to the database...........")

	var url = os.Getenv("url")

	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.HandleNilError(err)

	errs := DB.AutoMigrate(&model.Book{})
	helper.HandleNilError(errs)
}
