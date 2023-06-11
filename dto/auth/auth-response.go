package authdto

type LoginResponse struct {
	Id int `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Role string `json:"role"`
	Token    string `json:"token"`
}

type RegisterResponse struct {
	Email    string `json:"email"`
	Password string `json:"pasword"`
}