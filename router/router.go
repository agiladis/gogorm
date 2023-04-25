package router

import (
	"gogorm/controllers"
	"gogorm/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "gogorm/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	USERNAME = "admin"
	PASSWORD = "123456"
)

func StartServer(db *gorm.DB) *gin.Engine {

	userService := services.UserService{
		DB: db,
	}
	// userController will use userService to get data from db
	userController := controllers.UserController{
		UserService: &userService,
	}

	authController := controllers.AuthController{
		UserService: &userService,
	}

	app := gin.Default()

	app.POST("/login", authController.Login)

	authorized := app.Group("/users")

	authorized.GET("/", userController.GetAllUser)
	authorized.POST("", userController.CreateUser)
	authorized.GET("/:id", userController.GetUserById)
	authorized.PUT("/:id", userController.UpdateUserById)
	authorized.DELETE("/:id", userController.DeleteUserById)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return app
}
