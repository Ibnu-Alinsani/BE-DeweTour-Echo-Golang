package tripdto

import "dumbmerch/models"

type TripResponse struct {
	Title          string                 `json:"title"`
	CountryId      int                    `json:"country_id"`
	Country        models.CountryResponse `json:"country"`
	Accomodation   string                 `json:"accomodation"`
	Transportation string                 `json:"transportation"`
	Eat            string                 `json:"eat" form:"eat"`
	Day            string                 `json:"day"`
	Night          string                 `json:"night"`
	DateTrip       string                 `json:"date_trip"`
	Price          int                    `json:"price"`
	Quota          int                    `json:"quota"`
	Description    string                 `json:"description"`
	Image          string                 `json:"image"`
}