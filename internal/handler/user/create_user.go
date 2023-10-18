package user

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/service/user"
)

// UserRequest matches JSON request to create a user
type UserRequest struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func isValid(req UserRequest) error {
	// trim space
	req.Email = strings.TrimSpace(req.Email)

	// check if email exists
	if req.Email == "" {
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
		var req UserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return handler.HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "Invalid JSON request",
			}
		}

		if err := isValid(req); err != nil {
			return err
		}

		if _, err := h.userService.CreateUser(r.Context(), user.UserInput{
			Email: req.Email,
		}); err != nil {
			return handler.ConvertErr(err)
		}

		json.NewEncoder(w).Encode(handler.SuccessResponse{
			Success: true,
		})

		return nil
	})
}
