package DTOs

type LoginRequestDTO struct {
	Email string `json:"email" binding:"required,email"`
	Pwd   string `json:"pwd" binding:"required,min=8"`
}
