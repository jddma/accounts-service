package config

import (
	"cine-accounts/app/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DefinePaths(router *gin.RouterGroup) {
	router.POST("/login", controllers.Login)
	router.POST("/user", controllers.CreateNewUser)
}

func SetCors(router *gin.RouterGroup) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
}

func SetLogger(router *gin.RouterGroup) {
	router.Use(gin.Logger())
}
