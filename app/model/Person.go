package model

type Person struct {
	Id             int
	Names          string
	LastNames      string
	DocumentType   string
	DocumentNumber string
	BirthDate      string
}

func NewPerson(names string, lastNames string, documentType string, documentNumber string, birthDate string) *Person {
	return &Person{Names: names, LastNames: lastNames, DocumentType: documentType, DocumentNumber: documentNumber, BirthDate: birthDate}
}
