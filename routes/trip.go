package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func TripRoutes(e *echo.Group) {
	tripRepository := repository.RepositoryTrip(mysql.DB)
	h := handlers.HandlerTrip(tripRepository)

	e.GET("/trips", h.GetAllTrip)
	e.GET("/trip/:id", h.GetTripById)
	e.POST("/add-trip", middleware.Auth(middleware.UploadFile(h.CreateTrip)))
	e.DELETE("/delete-trip/:id", middleware.Auth(h.DeleteTrip))
	e.PATCH("/edit-trip/:id", middleware.Auth(middleware.UploadFile(h.EditTrip)))
}