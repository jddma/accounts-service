package DTOs

type LoginRequestDTO struct {
	Email string `json:"Email"`
	Pwd   string `json:"Pwd"`
}
