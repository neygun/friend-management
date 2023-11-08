package relationship

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/service/relationship"
	"github.com/neygun/friend-management/pkg/util"
)

// GetCommonFriendsRequest matches JSON request to retrieve the common friends list between two email addresses
type GetCommonFriendsRequest struct {
	Friends []string `json:"friends"`
}

func (req *GetCommonFriendsRequest) isValid() error {
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

// GetCommonFriends handles requests to retrieve the common friends list between two email addresses
func (h Handler) GetCommonFriends() http.HandlerFunc {
	return handler.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		var req GetCommonFriendsRequest
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

		// Get common friends
		commonFriends, count, err := h.relationshipService.GetCommonFriends(r.Context(), relationship.GetCommonFriendsInput{Friends: req.Friends})
		if err != nil {
			return handler.ConvertError(err)
		}

		// Success
		json.NewEncoder(w).Encode(handler.GetFriendsSuccess{
			Success: true,
			Friends: commonFriends,
			Count:   count,
		})

		return nil
	})
}
