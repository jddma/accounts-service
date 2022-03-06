package controllers

import (
	"accounts-service/app/DTOs"
	"accounts-service/app/business"
	"accounts-service/app/util"
	"github.com/gin-gonic/gin"
	"log"
)

func LoginController(ctx *gin.Context) {
	loginRequestDTO := new(DTOs.LoginRequestDTO)
	err := ctx.ShouldBindJSON(loginRequestDTO)
	if util.HandleError(ctx, err) {
		return
	}

	log.Printf("Usuario solicitado -> %s", loginRequestDTO.Email)
	business.LoginBusiness(loginRequestDTO, ctx)
}
