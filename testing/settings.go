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
