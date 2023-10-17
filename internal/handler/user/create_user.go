package user

import (
	"encoding/json"
	"net/http"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/model"
)

func isValid(input model.User) error {
	// check if email exists
	if input.Email == "" {
		return handler.HandlerErr{
			Code:        http.StatusBadRequest,
			Description: "Missing email field",
		}
	}
	return nil
}

// CreateUser handles requests to create a user
func (h Handler) CreateUser() http.HandlerFunc {
	return handler.ErrHandler(func(w http.ResponseWriter, r *http.Request) error {
		var input model.User
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return handler.HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "Invalid JSON request",
			}
		}

		if err := isValid(input); err != nil {
			return err
		}

		if _, err := h.userService.CreateUser(r.Context(), input); err != nil {
			return handler.ConvertErr(err)
		}

		json.NewEncoder(w).Encode(handler.SuccessResponse{
			Success: true,
		})

		return nil
	})
}
