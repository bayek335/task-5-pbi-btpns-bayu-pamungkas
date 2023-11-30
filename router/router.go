package router

import (
	"log"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/controllers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/middlewares"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/models"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {
	conf := app.GetConfig()
	router := gin.Default()

	usersModel := models.NewUser(db)
	usersService := services.NewUser(usersModel)
	userController := controllers.NewUser(usersService)
	authController := controllers.NewAuth(usersService)

	v1 := router.Group("/v1")
	users := v1.Group("/users")

	users.POST("/register", authController.Register)
	users.POST("/login", authController.Login)

	users.Use(middlewares.AuthJwtToken())
	users.PUT("/:id", userController.UpdateUser)
	users.DELETE("/:id", userController.DeleteUser)

	if err := router.Run(conf.ServerHost + ":" + conf.ServerPort); err != nil {
		log.Fatal(err.Error())
	}
}
