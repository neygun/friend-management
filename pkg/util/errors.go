package util

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
	pkgerrors "github.com/pkg/errors"
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
	var pgErr *pq.Error

	if errors.As(pkgerrors.Cause(err), &pgErr) {
		return string(pgErr.Code) == pe.Code
	}

	return false
}
