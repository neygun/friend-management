package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/jwtauth/v5"
	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/service/user"
	"github.com/neygun/friend-management/pkg/util"
)

// LoginRequest matches JSON request to logout
type LogoutRequest struct {
	Email string `json:"email"`
}

func (req *LogoutRequest) isValid() error {
	// trim space
	req.Email = strings.TrimSpace(req.Email)

	// check if email exists
	if req.Email == "" {
		return handler.HandlerError{
			Code:        http.StatusBadRequest,
			Description: "Missing email field",
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

// Logout handles requests to logout
func (h Handler) Logout() http.HandlerFunc {
	return handler.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		var req LogoutRequest
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

		// Logout
		tokenString := jwtauth.TokenFromHeader(r)
		if err := h.userService.Logout(r.Context(), user.LogoutInput{
			Email: req.Email,
			Token: tokenString,
		}); err != nil {
			return handler.ConvertError(err)
		}

		// Success
		json.NewEncoder(w).Encode(handler.SuccessResponse{
			Success: true,
		})

		return nil
	})
}
