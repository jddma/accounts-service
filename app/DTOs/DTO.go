package DTOs

type DTO interface{}

type DTOWithPwd interface {
	SetPwdForLog(string) CreateNewUserRequestDTO
}
