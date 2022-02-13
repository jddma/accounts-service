package app

import (
	"cine-accounts/app/config"
	"github.com/gin-gonic/gin"
	"os"
)

func Run() {
	r := gin.Default()

	g := r.Group("/api/v1")
	config.DefinePaths(g)

	r.Run(os.Getenv("port"))
}
