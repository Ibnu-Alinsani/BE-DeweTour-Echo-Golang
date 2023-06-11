package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	routes "dumbmerch/routes"
	"fmt"
	"os"

	// "net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {

	godotenv.Load()

	var e = echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	mysql.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")

	PORT := os.Getenv("PORT")

	fmt.Println("Running on port " + PORT)
	e.Logger.Fatal(e.Start(":" + PORT))
}

