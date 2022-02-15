package business

import (
	"cine-accounts/app/DTOs"
	"cine-accounts/app/repository"
	"cine-accounts/app/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginBusiness(loginRequestDTO *DTOs.LoginRequestDTO, ctx *gin.Context) {
	user, err := repository.FindByEmail(loginRequestDTO.Email)
	if validateConnectionWithDatabase(ctx, err) {
		return
	}
	if user.Id == 0 {
		log.Printf("No se ha encontrado el usuario -> %s", loginRequestDTO.Email)
		ctx.Status(http.StatusNotFound)
	}

	if util.ValidatePwd(loginRequestDTO.Pwd, user.Pwd) {
		log.Printf("Credenciales incorrectas para usuario -> %s", loginRequestDTO.Email)
		ctx.Status(http.StatusUnauthorized)
	} else {
		log.Printf("Acceso concedido para usuario -> %s", loginRequestDTO.Email)
		//To do Token
		ctx.Status(http.StatusOK)
	}
}
