package controllers

import (
	"accounts-service/app/business"
	"github.com/gin-gonic/gin"
	"log"
)

func ConsultRolesController(ctx *gin.Context) {
	log.Printf("Consulta de roles")
	business.ConsultRolesBusiness(ctx)
}
