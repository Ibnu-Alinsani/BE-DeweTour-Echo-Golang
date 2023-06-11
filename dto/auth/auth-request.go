package authdto

type AuthRequest struct {
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone string 	`json:"phone"`
	Address string `json:"address"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}