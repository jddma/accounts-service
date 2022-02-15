package controllers

import (
	"cine-accounts/app/DTOs"
	"cine-accounts/app/business"
	"cine-accounts/app/util"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateNewUser(ctx *gin.Context) {
	createNewUserRequestDTO := new(DTOs.CreateNewUserRequestDTO)
	err := ctx.Bind(createNewUserRequestDTO)
	handleError(ctx, err)

	log.Printf("Objeto request CreateNewUser -> %s", util.DTOWithPwdToJson(createNewUserRequestDTO))
	business.CreateNewUserBusiness(createNewUserRequestDTO, ctx)
}
