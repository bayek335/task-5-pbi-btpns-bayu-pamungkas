package database

import (
	"fmt"
	"log"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	conf := app.GetConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", conf.DatabaseHost, conf.DatabaseUser, conf.DatabasePassword, conf.DatabaseName, conf.DatabasePort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection to database error!")
	}

	return db
}
