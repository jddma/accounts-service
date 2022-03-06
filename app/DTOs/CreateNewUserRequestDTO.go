package DTOs

import "time"

type CreateNewUserRequestDTO struct {
	Names          string    `json:"names" binding:"required"`
	Lastnames      string    `json:"lastnames" binding:"required"`
	DocumentType   string    `json:"documentType" binding:"oneof=CC CE"`
	DocumentNumber string    `json:"documentNumber" binding:"required,numeric"`
	BirthDate      time.Time `json:"birthDate" binding:"required" time_format:"2006-01-02"`
	IdRole         int       `json:"idRole" binding:"required"`
	Email          string    `json:"email" binding:"required,email"`
	Pwd            string    `json:"pwd" binding:"required,min=8"`
	Phone          string    `json:"phone" binding:"required,numeric"`
}

func (c CreateNewUserRequestDTO) SetPwdForLog(pwd string) CreateNewUserRequestDTO {
	c.Pwd = pwd
	return c
}
