package repository

import (
	"database/sql"

	"github.com/sony/sonyflake"
)

type RelationshipRepository interface {
}

type relationshipRepository struct {
	db    *sql.DB
	idsnf *sonyflake.Sonyflake
}
