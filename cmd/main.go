package main

import (
	"log"
	"pet-info-backend-docker/infraestructure/handler"

	"pet-info-backend-docker/pkg/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	handler.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
