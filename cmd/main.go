package main

import (
	"database/sql"
	"log"
	"pet-info-backend-docker/infraestructure/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

// InitDB inicializa la conexión a la base de datos
func InitDB() (*sql.DB, error) {
	dbUser := "tu_usuario"
	dbPassword := "tu_contraseña"
	dbName := "mi_base_de_datos"
	dbHost := "db"
	connStr := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	handler.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
