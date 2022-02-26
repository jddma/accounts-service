package DTOs

type LoginRequestDTO struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}
