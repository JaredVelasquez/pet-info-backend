package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	var db *sql.DB
	var err error

	fmt.Printf("Connecting to the database at %s...\n", dbHost)
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("Connected to the database successfully!")

				// Ejecutar el script de migración
				if err := runMigration(db); err != nil {
					return nil, err
				}

				return db, nil
			}
		}
		fmt.Printf("Could not connect to the database: %v. Retrying...\n", err)
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to the database after multiple attempts: %v", err)
}

// runMigration ejecuta el script de migración para crear las tablas si no existen
func runMigration(db *sql.DB) error {
	file, err := ioutil.ReadFile("migration/users.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(file))
	return err
}
