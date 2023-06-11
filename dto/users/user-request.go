package userdto

type CreateUser struct {
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}

type UpdateUser struct {
	FullName string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	No_Handphone string `json:"no_handphone" form:"no_handphone"`
	Address string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}