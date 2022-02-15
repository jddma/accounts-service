package controllers

import (
	"cine-accounts/app/DTOs"
	"cine-accounts/app/business"
	"github.com/gin-gonic/gin"
	"log"
)

func Login(ctx *gin.Context) {
	loginRequestDTO := new(DTOs.LoginRequestDTO)
	err := ctx.Bind(loginRequestDTO)
	handleError(ctx, err)

	log.Printf("Usuario solicitado -> %s", loginRequestDTO.Email)
	business.LoginBusiness(loginRequestDTO, ctx)
}
