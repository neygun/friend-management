package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/neygun/friend-management/internal/handler"
)

var (
	errUnauthorized        = errors.New("unauthorized")
	errInternalServerError = errors.New("internal server error")
)

func errorHandler(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(handler.Response{
		Code:        code,
		Description: err.Error(),
	})
}
