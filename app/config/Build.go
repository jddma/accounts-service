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

	/*router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "uuid"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))*/
}
