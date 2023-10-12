package util

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Init() (*sql.DB, error) {
	connectionStr := "postgres://postgres:postgres@localhost:5432/fm-pg?sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	return db, err
}