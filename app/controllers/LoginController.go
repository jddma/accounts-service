package controllers

import (
	"accounts-service/app/DTOs"
	"accounts-service/app/business"
	"github.com/gin-gonic/gin"
	"log"
)

func LoginController(ctx *gin.Context) {
	loginRequestDTO := new(DTOs.LoginRequestDTO)
	err := ctx.Bind(loginRequestDTO)
	handleError(ctx, err)

	log.Printf("Usuario solicitado -> %s", loginRequestDTO.Email)
	business.LoginBusiness(loginRequestDTO, ctx)
}
