package model

import "time"

type Person struct {
	Id             int
	Names          string
	LastNames      string
	DocumentType   string
	DocumentNumber string
	BirthDate      time.Time
}

func NewPerson(names string, lastNames string, documentType string, documentNumber string, birthDate time.Time) *Person {
	return &Person{Names: names, LastNames: lastNames, DocumentType: documentType, DocumentNumber: documentNumber,
		BirthDate: birthDate}
}
