package DTOs

type GenerateTokenDTO struct {
	IdUser int `json:"idUser"`
	IdRole   int `json:"idRole"`
}

func NewGenerateTokenDTO(idUser int, idRole int) *GenerateTokenDTO {
	return &GenerateTokenDTO{IdUser: idUser, IdRole: idRole}
}
