package repository

import "cine-accounts/app/model"

func FindByEmail(email string) model.User {
	db := openConnection()
	var user model.User

	db.Where("Email = ?", email).First(&user)
	return user
}
