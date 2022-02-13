package business

import (
	"cine-accounts/app/DTO"
	"cine-accounts/app/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginBusiness(loginRequestDTO *DTO.LoginRequestDTO) int {
	user := repository.FindByEmail(loginRequestDTO.Email)
	if user.IdUser == 0 {
		return http.StatusNotFound
	}

	pwdValidate := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(loginRequestDTO.Pwd))
	if pwdValidate != nil {
		return http.StatusUnauthorized
	} else {
		return http.StatusOK
	}
	return user.IdUser
}
