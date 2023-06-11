package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func TransactionRoute(e *echo.Group) {
	transRepo := repository.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransactions(transRepo)

	e.GET("/transactions", h.GetAllTransactions)
	e.GET("/transaction/:id", h.GetTransactionById)
	e.POST("/add-transaction", middleware.Auth(h.CreateTransaction))
	// e.PATCH("/edit-transaction/:id", middleware.Auth(h.EditTransaction))
	e.POST("/notification", h.Notification)
	// e.PATCH("/edit-transactionStatus/:id", middleware.Auth(h.EditTransaction))
	// e.DELETE("/delete-transaction/:id", h.DeleteTransaction)
}