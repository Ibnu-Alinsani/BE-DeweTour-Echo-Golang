package transactiondto

type CreateTransactions struct {
	CounterQty int    `json:"counterQty" form:"counterQty"`
	Total      int    `json:"total" form:"total"`
	Status     string `json:"status" form:"status"`
	Attachment string `json:"attachment" form:"attachment"`
	TripId     int    `json:"tripId" form:"tripId"`
	UserId     int    `json:"user_id"`
}

type UpdateTransactions struct {
	CounterQty int    `json:"counterQty"`
	Total      int    `json:"total"`
	Status     string `json:"status"`
	Attachment string `json:"attachment"`
	TripId     int    `json:"tripId"`
}