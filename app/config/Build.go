package config

import (
	"accounts-service/app/controllers"
	"github.com/gin-gonic/gin"
)

func DefinePaths(router *gin.RouterGroup) {
	router.POST("/login", controllers.LoginController)
	router.POST("/user", controllers.CreateNewUserController)

	router.GET("/role", controllers.ConsultRolesController)
}

func SetLogger(router *gin.RouterGroup) {
	router.Use(gin.Logger())
}
