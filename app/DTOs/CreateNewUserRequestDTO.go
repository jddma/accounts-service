package DTOs

type CreateNewUserRequestDTO struct {
	Names          string `json:"names"`
	Lastnames      string `json:"lastnames"`
	DocumentType   string `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
	BirthDate      string `json:"birthDate"`
	IdRole         int    `json:"idRole"`
	Email          string `json:"email"`
	Pwd            string `json:"pwd"`
	Phone          string `json:"phone"`
}

func (c CreateNewUserRequestDTO) SetPwdForLog(pwd string) CreateNewUserRequestDTO {
	c.Pwd = pwd
	return c
}
