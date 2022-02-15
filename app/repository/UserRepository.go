package repository

import (
	"cine-accounts/app/model"
)

func FindByEmail(email string) (model.User, error) {
	db, err := openConnection()
	var user model.User
	db.Where("Email = ?", email).First(&user)
	return user, err
}

func FindByPersonAndRole(idPerson int, idRole int) (*model.User, error) {
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
