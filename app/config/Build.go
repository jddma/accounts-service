package config

import (
	"accounts-service/app/controllers"
	"accounts-service/app/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DefinePaths(router *gin.RouterGroup) {
	router.POST("/login", middleware.ValidateHeaders(), controllers.LoginController)
	router.POST("/user", middleware.ValidateAuthorization(), controllers.CreateNewUserController)

	router.GET("/role", middleware.ValidateAuthorization(), controllers.ConsultRolesController)
}

func SetLogger(router *gin.RouterGroup) {
	router.Use(gin.Logger())
}

func GetCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AddAllowHeaders("uuid", "Authorization")
	config.AllowAllOrigins = true
	return cors.New(config)
}
