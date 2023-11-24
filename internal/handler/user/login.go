package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/service/user"
	"github.com/neygun/friend-management/pkg/util"
)

// LoginRequest matches JSON request to login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginRequest) isValid() error {
	// trim space
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	// check if email exists
	if req.Email == "" {
		return handler.HandlerError{
			Code:        http.StatusBadRequest,
			Description: "Missing email field",
		}
	}

	// check if password exists
	if req.Password == "" {
		return handler.HandlerError{
			Code:        http.StatusBadRequest,
			Description: "Missing password field",
		}
	}

	// check if the email is valid
	if !util.IsEmail(req.Email) {
		return handler.HandlerError{
			Code:        http.StatusBadRequest,
			Description: "Invalid email",
		}
	}

	return nil
}

// Login handles requests to login
func (h Handler) Login() http.HandlerFunc {
	return handler.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		var req LoginRequest
		// Parse request body
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return handler.HandlerError{
				Code:        http.StatusBadRequest,
				Description: "Invalid JSON request",
			}
		}

		// Validate request
		if err := req.isValid(); err != nil {
			return err
		}

		// Login
		token, err := h.userService.Login(r.Context(), user.LoginInput{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			return handler.ConvertError(err)
		}

		// Success
		json.NewEncoder(w).Encode(LoginSuccess{
			Success: true,
			Token:   token,
		})

		return nil
	})
}
