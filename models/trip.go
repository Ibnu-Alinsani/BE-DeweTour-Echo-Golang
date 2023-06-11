package models

import "time"

type Trip struct {
	Id             int             `json:"id"`
	Title          string          `json:"title" form:"title" validate:"required" gorm:"varchar(255)"`
	CountryId      int             `json:"country_id" validate:"required" form:"country_id"`
	Country        CountryResponse `jsonssss:"country" form:"country" gorm:"foreignKey: CountryId;"`
	Accomodation   string          `json:"accomodation" form:"accomodation" validate:"required" gorm:"varchar(255)"`
	Transportation string          `json:"transportation" form:"transportation" validate:"required" gorm:"varchar(255)"`
	Eat            string          `json:"eat" form:"eat" gorm:"varchar(255)"`
	Day            string          `json:"day" form:"day" validate:"required" gorm:"varchar(255)"`
	Night          string          `json:"night" form:"night" validate:"required" gorm:"varchar(255)"`
	DateTrip       string          `json:"date_trip" form:"date_trip" validate:"required"`
	Price          int             `json:"price" form:"price" validate:"required"`
	Quota          int             `json:"quota" form:"quota" validate:"required"`
	CurrentQuota   int 		       `json:"current_quota"`
	Description    string          `json:"description" form:"description" validate:"required"`
	Image          string		   `json:"image" form:"image" validate:"required"`
	CreatedAt      time.Time       `json:"-"`
	UpdatedAt	   time.Time	   `json:"-"`
}

type TripResponse struct {
	Id             int             `json:"id"`
	Title          string          `json:"title"`
	CountryId      int             `json:"country_id"`
	Country        CountryResponse `json:"country"`
	Accomodation   string          `json:"accomodation"`
	Transportation string          `json:"transportation"`
	Eat            string          `json:"eat"`
	Day            string          `json:"day"`
	Night          string          `json:"night"`
	DateTrip       string          `json:"date_trip"`
	Price          int             `json:"price"`
	Quota          int             `json:"quota"`
	Description    string          `json:"description"`
	Image          string          `json:"image"`
}

func (TripResponse) TableName() string {
	return "trips"
}