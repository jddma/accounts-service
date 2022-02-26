package business

import (
	"accounts-service/app/repository"
	"accounts-service/app/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ConsultRolesBusiness(ctx *gin.Context) {
	roles, err := repository.FindRoles()
	if validateConnectionWithDatabase(ctx, err) {
		return
	}
	log.Printf("Roles obtenidos de la DB -> %s", util.DtoToJson(roles))
	ctx.JSON(http.StatusOK, roles)
}
