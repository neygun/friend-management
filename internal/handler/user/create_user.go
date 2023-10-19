package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/pkg/util"
)

// UserRequest matches JSON request to create a user
type UserRequest struct {
	Email string `json:"email"`
}

func (req *UserRequest) isValid() error {
	// trim space
	req.Email = strings.TrimSpace(req.Email)

	// check if email exists
	if req.Email == "" {
		return handler.HandlerErr{
			Code:        http.StatusBadRequest,
			Description: "Missing email field",
		}
	}

	// check if the email is valid
	if !util.IsEmail(req.Email) {
		return handler.HandlerErr{
			Code:        http.StatusBadRequest,
			Description: "Invalid email",
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

		if err := req.isValid(); err != nil {
			return err
		}

		if _, err := h.userService.CreateUser(r.Context(), model.User{
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
