package relationship

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/service/relationship"
	"github.com/neygun/friend-management/pkg/util"
)

// FriendConnectionRequest matches JSON request to create a friend connection
type FriendConnectionRequest struct {
	Friends []string `json:"friends"`
}

func (req *FriendConnectionRequest) isValid() error {
	// trim space
	for i, v := range req.Friends {
		req.Friends[i] = strings.TrimSpace(v)
	}

	// check valid emails
	for _, v := range req.Friends {
		if !util.IsEmail(v) {
			return handler.HandlerError{
				Code:        http.StatusBadRequest,
				Description: "Invalid email",
			}
		}
	}

	// check if number of emails = 2
	if len(req.Friends) != 2 {
		return handler.HandlerError{
			Code:        http.StatusBadRequest,
			Description: "The number of emails must be 2",
		}
	}

	// check if the emails are the same
	if req.Friends[0] == req.Friends[1] {
		return handler.HandlerError{
			Code:        http.StatusBadRequest,
			Description: "The emails are the same",
		}
	}
	return nil
}

// CreateFriendConnection handles requests to create friend connection
func (h Handler) CreateFriendConnection() http.HandlerFunc {
	return handler.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		var req FriendConnectionRequest
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

		// Create friend connection
		if _, err := h.relationshipService.CreateFriendConnection(r.Context(), relationship.FriendConnectionInput{
			Friends: req.Friends,
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
