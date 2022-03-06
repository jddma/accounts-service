package business

import (
	"accounts-service/app/DTOs"
	"accounts-service/app/repository"
	"accounts-service/app/requestCatalog"
	"accounts-service/app/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginBusiness(loginRequestDTO *DTOs.LoginRequestDTO, headersDTO *DTOs.AnonymousHeadersDTO, ctx *gin.Context) {
	user, err := repository.FindUserByEmail(loginRequestDTO.Email)
	if validateConnectionWithDatabase(ctx, err) {
		return
	}
	if user.Id == 0 {
		log.Printf("No se ha encontrado el usuario -> %s", loginRequestDTO.Email)
		ctx.Status(http.StatusNotFound)
		return
	}
	log.Printf("Usuario obtenido de la DB -> %s", util.DtoToJson(user))
	if util.ValidatePwd(loginRequestDTO.Pwd, user.Pwd) {
		log.Printf("Acceso concedido para usuario -> %s", loginRequestDTO.Email)
		tokenResponseDTO, nil := requestCatalog.ConsumeJwtGenerator(user.Id, user.IdRole, headersDTO, ctx)
		if err != nil {
			return
		}

		ctx.JSON(http.StatusOK, tokenResponseDTO)
	} else {
		log.Printf("Credenciales incorrectas para usuario -> %s", loginRequestDTO.Email)
		ctx.Status(http.StatusUnauthorized)
	}
}
