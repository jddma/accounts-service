package controllers

import (
	"accounts-service/app/DTOs"
	"accounts-service/app/business"
	"accounts-service/app/util"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateNewUserController(ctx *gin.Context) {
	createNewUserRequestDTO := new(DTOs.CreateNewUserRequestDTO)
	err := ctx.Bind(createNewUserRequestDTO)
	handleError(ctx, err)

	log.Printf("Objeto request CreateNewUser -> %s", util.DTOWithPwdToJson(createNewUserRequestDTO))
	business.CreateNewUserBusiness(createNewUserRequestDTO, ctx)
}
