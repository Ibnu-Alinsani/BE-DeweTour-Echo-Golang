package routes

import (
	// "dumbmerch/handlers"
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repository.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.GetAllUser)
	e.GET("/user/:id", h.GetUserById)
	// e.POST("/create-account", h.CreateUser)
	e.PATCH("/update-user/:id", middleware.Auth(middleware.UploadFile(h.UpdateUser)))
	e.DELETE("/delete-user/:id", middleware.Auth(h.DeleteUser))
}