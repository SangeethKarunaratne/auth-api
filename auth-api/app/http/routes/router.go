package routes

import (
	"auth-api/app/container"
	"auth-api/app/http/controllers"
	"auth-api/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine, ctr *container.Container) {

	userController := controllers.NewUserController(ctr)

	server.POST("/login", userController.Login)
	server.POST("/register", userController.Register)

	api := server.Group("/api")
	api.Use(middleware.AuthMiddleware(ctr))
	api.GET("/users", userController.GetUsers)
}
