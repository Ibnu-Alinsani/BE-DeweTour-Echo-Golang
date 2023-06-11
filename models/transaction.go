package models

import "time"

type Transaction struct {
	Id         int          `json:"id" gorm:"primaryKey;auto_increment"`
	CounterQty int          `json:"counterQty" form:"counterQty"`
	Total      int          `json:"total" form:"form"`
	Status     string       `json:"status" form:"status"`
	Attachment string       `json:"attachment" form:"attachment"`
	UserId     int          `json:"user_id"`
	User       UserResponse `json:"user" gorm:"foreignKey:UserId"`
	TripId     int          `json:"tripId" form:"tripId"`
	Trip       TripResponse `json:"trip" gorm:"foreignKey:TripId"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"-"`
}

type TransactionResponse struct {
	Id         int    `json:"id" gorm:"primaryKey;auto_increment"`
	CounterQty int    `json:"counterQty" form:"counterQty"`
	Total      int    `json:"total" form:"form"`
	Status     string `json:"status" form:"status"`
	Attachment string `json:"attachment" form:"attachment"`
	UserId     int    `json:"user_id"`
	TripId     int    `json:"tripId" form:"tripId"`
	Trip TripResponse `json:"trip"`
	CreatedAt  time.Time `json:"created_at"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}