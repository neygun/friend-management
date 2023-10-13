package user

import (
	"encoding/json"
	"net/http"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/model"
)

func (h UserHandler) CreateUser() http.HandlerFunc {
	return handler.ErrHandler(func(w http.ResponseWriter, r *http.Request) error {
		var input model.User
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return handler.HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "Invalid JSON request",
			}
		}

		// check if email exists
		if input.Email == "" {
			return handler.HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "Missing email field",
			}
		}

		if _, err := h.userService.CreateUser(r.Context(), input); err != nil {
			return handler.ConvertErr(err)
		}

		json.NewEncoder(w).Encode(model.SuccessResponse{
			Success: true,
		})

		return nil
	})
}
