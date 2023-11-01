package util

import (
	"fmt"

	friendErr "github.com/friendsofgo/errors"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgError struct {
	Code string
	Msg  string
}

func (pe pgError) Error() string {
	return fmt.Sprintf("%s: %s", pe.Code, pe.Msg)
}

// Add more errors if needed from
// https://www.postgresql.org/docs/14/errcodes-appendix.html
var (
	// UniqueViolation is a unique violation error
	UniqueViolation = pgError{Code: "23505", Msg: "duplicate key value violates unique constraint"}
	// CheckViolation is a check key violation error
	CheckViolation = pgError{Code: "23514", Msg: "check violation for the key"}
)

// Is checks if the error is a pgError
func (pe pgError) Is(err error) bool {
	var pgErr *pgconn.PgError

	if friendErr.As(friendErr.Cause(err), &pgErr) {
		return pgErr.Code == pe.Code
	}

	return false
}
