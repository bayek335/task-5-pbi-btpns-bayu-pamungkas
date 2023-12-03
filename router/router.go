package router

import (
	"log"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/controllers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/middlewares"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/models"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {
	conf := app.GetConfig()
	router := gin.Default()

	// static file
	router.Static("public", "./public")

	router.Use(middlewares.Cors())
	router.Use(cors.New(cors.Config{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"POST", "PUT", "GET", "DELETE", "OPTION"},
		AllowedHeaders: []string{"*"},
	}))
	// initiate session
	store := cookie.NewStore([]byte(conf.SessionKey))
	router.Use(sessions.Sessions("user_session", store))

	// user pattern
	usersModel := models.NewUser(db)
	usersService := services.NewUser(usersModel)
	userController := controllers.NewUser(usersService)
	authController := controllers.NewAuth(usersService)

	// photo pattern
	photoModel := models.NewPhoto(db)
	photoService := services.NewPhoto(photoModel)
	photoController := controllers.NewPhoto(photoService)
	v1 := router.Group("/v1")

	users := v1.Group("/users")
	photos := v1.Group("/photos")

	// user routes without auth
	users.POST("/register", authController.Register)
	users.POST("/login", authController.Login)
	// user routes with auth
	users.Use(middlewares.AuthJwtToken())
	users.GET("/profile", userController.FindUserByID)
	users.GET("/profile/:id", userController.FindUserByID)
	users.PUT("/:id", userController.UpdateUser)
	users.DELETE("/:id", userController.DeleteUser)

	// photo routes with auth
	photos.Use(middlewares.AuthJwtToken())
	photos.GET("", photoController.GetPhotos)
	photos.POST("", photoController.CreatePhoto)
	photos.POST("/", photoController.CreatePhoto)
	photos.PUT("/:id", photoController.UpdatePhoto)
	photos.DELETE("/:id", photoController.DeletePhoto)

	if err := router.Run(conf.ServerHost + ":" + conf.ServerPort); err != nil {
		log.Fatal(err.Error())
	}
}
