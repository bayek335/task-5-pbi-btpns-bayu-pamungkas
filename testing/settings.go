package testing

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func DatabaseConnection() *gorm.DB {

	dsn := "host=localhost user=postgres password=root dbname=pbi_btpns port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection to database error!")
	}

	return db
}

var User_ID string = "b2f53ff5-6389-4ee5-830d-42a8d73d6952"
