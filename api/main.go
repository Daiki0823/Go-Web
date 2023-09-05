package main

import (
	"api/handlers"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// health check
	e.GET("/api/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})
	// Routes
	e.GET("/api/todos", handlers.GetAllToDos)
	e.GET("/api//todo/:id", handlers.GetToDoByID)
	e.POST("/api//todo", handlers.AddToDo)

	// Determine the port based on the STAGE environment variable
	port := ":1323" // default to 1323
	if os.Getenv("STAGE") == "prod" {
		port = ":80"
	}
	// Start the server
	e.Start(port)
}
