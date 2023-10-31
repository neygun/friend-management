package util

import (
	"errors"

	"github.com/lib/pq"
)

func GetErrorCode(err error) string {
	pqErr := errors.Unwrap(errors.Unwrap(err)).(*pq.Error)
	return string(pqErr.Code)
}
