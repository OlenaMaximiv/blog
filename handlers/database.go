package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitializeDB initializes the database connection
func InitializeDB() error {
	connStr := "user=postgres dbname=mywebapp password=<your-password> sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	db = conn
	fmt.Println("Connected to the database")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Database connection closed")
	}
}
