package userdto

import "dumbmerch/models"

type UserResponse struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role string `json:"role"`
	Gender string `json:"gender"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Transaction []models.TransactionResponse `json:"transaction"`
}