package requestCatalog

import (
	"accounts-service/app/DTOs"
	"accounts-service/app/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func addHeader(r *http.Request) {
	r.Header.Add("Content-Type", "application/json")
}

func getAuthSrvURL(path string) string {
	URL := fmt.Sprintf("%s%s", os.Getenv("auth_srv"), path)
	return URL
}

func ConsumeJwtGenerator(idUser int, idRole int, ctx *gin.Context) (*DTOs.TokenDTO, error) {
	client := &http.Client{}
	generateTokenRequestDTO := DTOs.NewGenerateTokenDTO(idUser, idRole)
	log.Printf("Objeto request para authSrv GenerateJWT -> %s", util.DtoToJson(generateTokenRequestDTO))
	body := strings.NewReader(util.DtoToJson(generateTokenRequestDTO))
	request, _ := http.NewRequest("POST", getAuthSrvURL("/api/v1/token"), body)
	addHeader(request)

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error al consumir API JwtGenerator de authSrv -> %s", util.DtoToJson(err))
		ctx.JSON(http.StatusInternalServerError, util.ErrorMsg{Code: "EUX", Description: err.Error()})
		return nil, err
	}
	var tokenResponseDTO *DTOs.TokenDTO
	responseBytes, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseBytes, &tokenResponseDTO)
	log.Printf("Respuesta de authSrv -> %s", util.DtoToJson(tokenResponseDTO))
	return tokenResponseDTO, nil
}
