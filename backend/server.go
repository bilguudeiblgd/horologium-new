package main

import (
	"context"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func initializeDatabase() *pgx.Conn {
	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)

	}

	//defer conn.Close(context.Background())
	return conn
}

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	_ = initializeDatabase()
	// Define a simple GET endpoint
	e.GET("/ping", func(c *echo.Context) error {
		// Return JSON response
		return c.String(http.StatusOK, "pong!")
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
