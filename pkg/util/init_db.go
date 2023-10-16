package util

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// Init init a database connection
func Init() (*sql.DB, error) {
	connectionStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connectionStr)
	return db, err
}
