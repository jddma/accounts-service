package controllers

import (
	"cine-accounts/app/DTO"
	"cine-accounts/app/business"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	loginRequestDTO := new(DTO.LoginRequestDTO)
	err := ctx.Bind(loginRequestDTO)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	ctx.Status(business.LoginBusiness(loginRequestDTO))
}
