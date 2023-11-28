package router

import (
	"log"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {
	conf := app.GetConfig()
	router := gin.Default()

	if err := router.Run(conf.ServerHost + ":" + conf.ServerPort); err != nil {
		log.Fatal(err.Error())
	}
}
