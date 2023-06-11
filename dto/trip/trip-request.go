package tripdto

import "dumbmerch/models"

type CreateTrip struct {
	Title          string `json:"title" form:"title" validate:"required"`
	CountryId      int    `json:"country_id" validate:"required" form:"country_id"`
	Accomodation   string `json:"accomodation" form:"accomodation" validate:"required"`
	Transportation string `json:"transportation" form:"transportation" validate:"required"`
	Eat            string `json:"eat" form:"eat"`
	Day            string `json:"day" form:"day" validate:"required"`
	Night          string `json:"night" form:"night" validate:"required"`
	DateTrip       string `json:"date_trip" form:"date_trip" validate:"required"`
	Price          int    `json:"price" form:"price" validate:"required"`
	Quota          int    `json:"quota" form:"quota" validate:"required"`
	CurrentQuota int `json:"currentquota"`
	Description    string `json:"description" form:"description" validate:"required"`
	Image          string `json:"image" form:"image" validate:"required"`
}

type UpdateTrip struct {
	Title          string `json:"title" form:"title"`
	CountryId      int    `json:"country_id" form:"country_id"`
	Country        models.CountryResponse `json:"country"`
	Accomodation   string `json:"accomodation" form:"accomodation"`
	Transportation string `json:"transportation" form:"transportation"`
	Eat            string `json:"eat" form:"eat"`
	Day            string `json:"day" form:"day"`
	Night          string `json:"night" form:"night"`
	DateTrip       string `json:"date_trip" form:"date_trip"`
	Price          int    `json:"price" form:"price"`
	Quota          int    `json:"quota" form:"quota"`
	CurrentQuota int `json:"currentquota"`
	Description    string `json:"description" form:"description"`
	Image          string `json:"image" form:"image"`
}