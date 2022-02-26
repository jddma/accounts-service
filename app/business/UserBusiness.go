package business

import (
	"accounts-service/app/DTOs"
	"accounts-service/app/model"
	"accounts-service/app/repository"
	"accounts-service/app/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func registerPerson(createNewUserRequestDTO *DTOs.CreateNewUserRequestDTO) (*model.Person, error) {
	person := model.NewPerson(
		createNewUserRequestDTO.Names,
		createNewUserRequestDTO.Lastnames,
		createNewUserRequestDTO.DocumentType,
		createNewUserRequestDTO.DocumentNumber,
		createNewUserRequestDTO.BirthDate,
	)
	_, err := repository.CreatePerson(person)
	return person, err
}

func registerUser(createNewUserRequestDTO *DTOs.CreateNewUserRequestDTO, person *model.Person) (error, error, *model.User) {
	user := model.NewUser(
		createNewUserRequestDTO.IdRole,
		person.Id,
		createNewUserRequestDTO.Email,
		util.GeneratePwdHash(createNewUserRequestDTO.Pwd),
		createNewUserRequestDTO.Phone,
	)
	createErr, connectionErr := repository.CreateUser(user)
	return createErr, connectionErr, user
}

func CreateNewUserBusiness(createNewUserRequestDTO *DTOs.CreateNewUserRequestDTO, ctx *gin.Context) {
	user, err := repository.FindUserByEmail(createNewUserRequestDTO.Email)
	if validateConnectionWithDatabase(ctx, err) {
		return
	}
	if user.Id != 0 {
		log.Printf("Ya existe un usuario para el correo -> %s", createNewUserRequestDTO.Email)
		ctx.Status(http.StatusNotFound)
	}
	//Validar si ya existe la persona
	person, err := repository.FindPersonByDocumentNumber(createNewUserRequestDTO.DocumentNumber)
	if validateConnectionWithDatabase(ctx, err) {
		return
	}
	if person.Id == 0 {
		person, err = registerPerson(createNewUserRequestDTO)
		if validateConnectionWithDatabase(ctx, err) {
			return
		}
		log.Printf("Persona ingresada en portal -> %s", util.DtoToJson(person))

	} else {
		log.Printf("Persona consultada en portal para cración de nuevo usuario -> %s", util.DtoToJson(person))
		userValidate, err := repository.FindUserByPersonAndRole(person.Id, createNewUserRequestDTO.IdRole)
		if validateConnectionWithDatabase(ctx, err) {
			return
		}
		if userValidate.Id != 0 {
			log.Printf("Ya existe un usuario para el id_persona -> %s y id_role -> %s", person.Id,
				createNewUserRequestDTO.IdRole)
			ctx.JSON(http.StatusBadRequest, util.UserDuplicateError)
			return
		}
	}
	createErr, connectionErr, userDb := registerUser(createNewUserRequestDTO, person)
	if validateConnectionWithDatabase(ctx, connectionErr) {
		return
	}
	if createErr != nil {
		if strings.Contains(createErr.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusBadRequest, util.EmailDuplicateError)
			return
		}
	}
	log.Printf("Usuario creado en portal -> %s", util.DtoToJson(userDb))
	ctx.Status(http.StatusCreated)
}
