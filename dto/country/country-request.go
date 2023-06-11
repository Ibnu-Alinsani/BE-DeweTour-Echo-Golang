package countrydto

type CountryRequest struct {
	Name string `json:"name" form:"name"`
}
