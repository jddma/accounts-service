package controllers

import (
	"accounts-service/app/DTOs"
	"accounts-service/app/business"
	"github.com/gin-gonic/gin"
	"log"
)

func LoginController(ctx *gin.Context) {
	headersDTO, err := getAnonymousHeaders(ctx)
	if handleError(ctx, err) {
		return
	}

	loginRequestDTO := new(DTOs.LoginRequestDTO)
	err = ctx.ShouldBindJSON(loginRequestDTO)
	if handleError(ctx, err) {
		return
	}

	log.Printf("Usuario solicitado -> %s", loginRequestDTO.Email)
	business.LoginBusiness(loginRequestDTO, headersDTO, ctx)
}
