package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strconv"
)

func ValidatePwd(pwd string, pwdHash string) bool {
	pwdValidate := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(pwd))
	if pwdValidate == nil {
		return true
	} else {
		return false
	}
}

func GeneratePwdHash(pwd string) string {
	cost, err := strconv.Atoi(os.Getenv("pwd_cost"))
	if err != nil {
		log.Fatalf("No se configuro correctamente la variable pwd_cost")
	}
	pwdHash, _ := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	return string(pwdHash)
}
