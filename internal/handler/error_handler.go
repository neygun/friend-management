package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/neygun/friend-management/internal/service/relationship"
)

// HandlerErr represents errors that are not server errors
type HandlerErr struct {
	Code        int
	Description string
}

// Error implemented
func (e HandlerErr) Error() string {
	return e.Description
}

// ConvertErr converts service errors to handler errors
func ConvertErr(err error) error {
	switch err {
	case relationship.ErrUserNotFound:
		return HandlerErr{
			Code:        http.StatusBadRequest,
			Description: err.Error(),
		}
	case relationship.ErrFriendConnectionExists:
		return HandlerErr{
			Code:        http.StatusBadRequest,
			Description: err.Error(),
		}
	case relationship.ErrBlockExists:
		return HandlerErr{
			Code:        http.StatusBadRequest,
			Description: err.Error(),
		}
	default: // Unexpected error
		return err
	}
}

// ErrHandler handles errors from handler methods and responses JSON
func ErrHandler(handlerFunc func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(w, r); err != nil {
			herr, ok := err.(HandlerErr)
			if ok {
				w.WriteHeader(herr.Code)

				json.NewEncoder(w).Encode(Response{
					Code:        herr.Code,
					Description: herr.Description,
				})
				log.Println(err.Error())

				return
			}

			w.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(w).Encode(Response{
				Code:        http.StatusInternalServerError,
				Description: "Internal Server Error",
			})
			log.Println(err.Error())
		}
	}
}
