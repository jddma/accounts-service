package util

import (
	"accounts-service/app/DTOs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return true
	}
	return false
}

func GetAnonymousHeaders(ctx *gin.Context) (*DTOs.AnonymousHeadersDTO, error) {
	headersDTO := new(DTOs.AnonymousHeadersDTO)
	err := ctx.ShouldBindHeader(&headersDTO)
	return headersDTO, err
}

func GetHeadersWithIdentityDTO(ctx *gin.Context) (*DTOs.HeadersWithIdentityDTO, error) {
	headersDTO := new(DTOs.HeadersWithIdentityDTO)
	err := ctx.ShouldBindHeader(&headersDTO)
	return headersDTO, err
}
