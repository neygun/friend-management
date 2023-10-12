package relationship

import (
	"encoding/json"
	"net/http"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/service/relationship"
)

type HandlerErr struct {
	Code        int
	Description string
}

func (e HandlerErr) Error() string {
	return e.Description
}

func convertErr(err error) error {
	switch err {
	case relationship.ErrInvalidUsersLength:
		// Add more expected error
		return HandlerErr{
			Code:        http.StatusBadRequest,
			Description: err.Error(),
		}
	default: // Unexpected error
		return err
	}
}

func ErrHandler(handlerFunc func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(w, r); err != nil {
			herr, ok := err.(HandlerErr)
			if ok {
				w.WriteHeader(herr.Code)

				json.NewEncoder(w).Encode(model.Response{
					Code:        herr.Code,
					Description: herr.Description,
				})
				// log

				return
			}

			w.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(w).Encode(model.Response{
				Code:        http.StatusInternalServerError,
				Description: "Internal Server Error",
			})
		}
	}
}
