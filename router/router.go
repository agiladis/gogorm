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

func StartServer(db *gorm.DB) *gin.Engine {

	userService := services.UserService{
		DB: db,
	}
	// userController will use userService to get data from db
	userController := controllers.UserController{
		UserService: &userService,
	}

	app := gin.Default()
	app.GET("/users", userController.GetAllUser)
	app.POST("/users", userController.CreateUser)
	app.GET("/users/:id", userController.GetUserById)
	app.PUT("/users/:id", userController.UpdateUserById)
	app.DELETE("/users/:id", userController.DeleteUserById)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return app
}
