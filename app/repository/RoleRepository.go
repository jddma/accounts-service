package repository

import "accounts-service/app/model"

func FindRoles() ([]model.Role, error) {
	db, err := openConnection()
	var roles []model.Role
	db.Find(&roles)
	return roles, err
}
