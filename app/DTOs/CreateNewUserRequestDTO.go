package DTOs

type CreateNewUserRequestDTO struct {
	Names          string `json:"Names"`
	Lastnames      string `json:"Lastnames"`
	DocumentType   string `json:"DocumentType"`
	DocumentNumber string `json:"DocumentNumber"`
	BirthDate      string `json:"BirthDate"`
	IdRole         int    `json:"idRole"`
	Email          string `json:"Email"`
	Pwd            string `json:"Pwd"`
	Phone          string `json:"Phone"`
}

func (c CreateNewUserRequestDTO) SetPwdForLog(pwd string) CreateNewUserRequestDTO {
	c.Pwd = pwd
	return c
}
