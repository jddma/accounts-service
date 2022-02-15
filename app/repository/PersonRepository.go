package repository

import (
	"cine-accounts/app/model"
	"gorm.io/gorm"
)

func FindByDocumentNumber(documentNumber string) (*model.Person, error) {
	db, err := openConnection()
	var person model.Person
	db.Where("document_number = ?", documentNumber).First(&person)
	return &person, err
}

func CreatePerson(person *model.Person) (*gorm.DB, error) {
	db, err := openConnection()
	result := db.Create(&person)
	return result, err
}
