package models

type Country struct {
	Id int 	`json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type CountryResponse struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

func(CountryResponse) TableName() string {
	return "countries"
}