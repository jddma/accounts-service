package model

type User struct {
	Id       int
	IdRole   int
	IdPerson int
	Email    string
	Pwd      string
	Phone    string
}

func NewUser(idRole int, idPerson int, email string, pwd string, phone string) *User {
	return &User{IdRole: idRole, IdPerson: idPerson, Email: email, Pwd: pwd, Phone: phone}
}
