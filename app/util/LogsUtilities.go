package util

import (
	"accounts-service/app/DTOs"
	"encoding/json"
)

func DtoToJson(dto DTOs.DTO) string {
	jsonString, _ := json.Marshal(dto)
	return string(jsonString)
}

func DTOWithPwdToJson(dto DTOs.DTOWithPwd) string {
	dtoCopy := dto.SetPwdForLog("***")
	return DtoToJson(dtoCopy)
}
