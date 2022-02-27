package DTOs

type GenerateTokenDTO struct {
	IdUser int `json:"idUser"`
	Role   int `json:"role"`
}

func NewGenerateTokenDTO(idUser int, idRole int) *GenerateTokenDTO {
	return &GenerateTokenDTO{IdUser: idUser, Role: idRole}
}
