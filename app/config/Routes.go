package config

import (
	"cine-accounts/app/controllers"
	"github.com/gin-gonic/gin"
)

func DefinePaths(router *gin.RouterGroup) {
	router.POST("/login", controllers.Login)
}
