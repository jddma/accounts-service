package middleware

import (
	"accounts-service/app/requestCatalog"
	"accounts-service/app/util"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func ValidateAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Middleware de auth")
		headersDTO, err := util.GetHeadersWithIdentityDTO(ctx)
		if util.HandleError(ctx, err) {
			return
		}
		tokenArray := strings.Split(headersDTO.Token, "Bearer ")
		if len(tokenArray) != 2 {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("Header Authorization invalid"))
			return
		}
		token := tokenArray[1]
		authStatusCode, err := requestCatalog.ConsumeJWTValidator(token, ctx)
		if err != nil {
			ctx.Abort()
		}
		switch authStatusCode {
		case http.StatusOK:
			ctx.Set("token", token)
			ctx.Next()
			break
		case http.StatusUnauthorized:
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
			break
		default:
			ctx.Status(http.StatusInternalServerError)
			ctx.Abort()
			break
		}
	}
}
