package middleware

import (
	"accounts-service/app/util"
	"github.com/gin-gonic/gin"
	"log"
)

func ValidateHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Middleware de auth")
		_, err := util.GetAnonymousHeaders(ctx)
		if util.HandleError(ctx, err) {
			return
		}
		ctx.Next()
	}
}
