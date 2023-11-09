package util

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/friendsofgo/errors"
	_ "github.com/lib/pq"
)

// InitDB init a database connection
func InitDB() (*sql.DB, error) {
	connectionStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, errors.Errorf("db connect error: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Errorf("db ping error: %s", err.Error())
	}

	return db, nil
}
