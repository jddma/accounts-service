package config

import (
	"accounts-service/app/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DefinePaths(router *gin.RouterGroup) {
	router.POST("/login", controllers.LoginController)
	router.POST("/user", controllers.CreateNewUserController)

	router.GET("/role", controllers.ConsultRolesController)
}

func SetCors(router *gin.RouterGroup) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
}

func SetLogger(router *gin.RouterGroup) {
	router.Use(gin.Logger())
}
