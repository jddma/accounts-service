package requestCatalog

import (
	"accounts-service/app/DTOs"
	"accounts-service/app/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func addHeader(r *http.Request, headersDTO *DTOs.AnonymousHeadersDTO) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("uuid", headersDTO.Uuid)
}

func getAuthSrvURL(path string) string {
	URL := fmt.Sprintf("%s%s", os.Getenv("auth_srv"), path)
	return URL
}

func generateRequest(path string, method string, dto DTOs.DTO, auth bool, token string, ctx *gin.Context) (*http.Client, *http.Request) {
	client := &http.Client{}
	body := strings.NewReader(util.DtoToJson(dto))
	request, _ := http.NewRequest(method, getAuthSrvURL(path), body)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("uuid", ctx.GetHeader("uuid"))
	if auth {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	return client, request
}

func ConsumeJWTValidator(token string, ctx *gin.Context) (int, error) {
	client, request := generateRequest("/api/v1/validate-token", "POST", nil, true, token, ctx)
	response, err := client.Do(request)

	if err != nil {
		log.Printf("Error al consumir API KWT validator de authSrv -> %s", util.DtoToJson(err))
		ctx.JSON(http.StatusInternalServerError, util.ErrorMsg{Code: "EUX", Description: err.Error()})
		return 0, err
	}
	return response.StatusCode, nil
}

func ConsumeJwtGenerator(idUser int, idRole int, ctx *gin.Context) (*DTOs.TokenDTO, error) {
	generateTokenRequestDTO := DTOs.NewGenerateTokenDTO(idUser, idRole)
	client, request := generateRequest("/api/v1/token", "POST", generateTokenRequestDTO, false, "", ctx)

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error al consumir API JwtGenerator de authSrv -> %s", util.DtoToJson(err))
		ctx.JSON(http.StatusInternalServerError, util.ErrorMsg{Code: "EUX", Description: err.Error()})
		return nil, err
	}
	if response.StatusCode == http.StatusOK {
		var tokenResponseDTO *DTOs.TokenDTO
		responseBytes, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(responseBytes, &tokenResponseDTO)
		log.Printf("Respuesta de authSrv -> %s", util.DtoToJson(tokenResponseDTO))
		return tokenResponseDTO, nil
	} else {
		log.Printf("Error en la orquestacipn a authSrv")
		ctx.JSON(http.StatusInternalServerError, util.OrchestrationError)
		return nil, errors.New("Error en orquestacion con authSrv")
	}
}
