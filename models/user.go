package models

import "time"

type User struct {
	ID        int    `json:"id"`
	Image     string `json:"image"`
	FullName  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone 	  string `json:"phone"`
	Address   string `json:"address"`
	Gender string `json:"gender"`
	Role string `json:"role"`
	Transaction []TransactionResponse `json:"transaction"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserResponse struct {
	Id int 	`json:"id"`
	FullName string `json:"fullName"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	Gender string `json:"gender"` 
}

func (UserResponse) TableName() string {
	return "users"
}