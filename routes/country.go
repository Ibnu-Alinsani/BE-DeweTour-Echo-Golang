package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func CountryRoutes(e *echo.Group) {
	countryRepository := repository.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(countryRepository)

	e.GET("/country", h.GetAllCountry)
	e.GET("/country/:id", h.GetCountryById)
	e.POST("/add-country", middleware.Auth(h.AddCountry))
	e.PATCH("/edit-country/:id", middleware.Auth(h.EditCountry))
	e.DELETE("/delete-country/:id", middleware.Auth(h.DeleteCountry))
}