package repository

import (
	"accounts-service/app/model"
)

func FindUserByEmail(email string) (model.User, error) {
	db, err := openConnection()
	var user model.User
	db.Where("Email = ?", email).First(&user)
	return user, err
}

func FindUserByPersonAndRole(idPerson int, idRole int) (*model.User, error) {
	db, err := openConnection()
	var user model.User
	db.Where("id_person = ? AND id_role = ?", idPerson, idRole).First(&user)
	return &user, err
}

func CreateUser(user *model.User) (error, error) {
	db, err := openConnection()
	result := db.Create(&user)
	return result.Error, err

}
