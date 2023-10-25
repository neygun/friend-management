package relationship

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/service/relationship"
	"github.com/neygun/friend-management/pkg/util"
)

// GetFriendsRequest matches JSON request to retrieve the friends list for an email address
type GetFriendsRequest struct {
	Email string `json:"email"`
}

func (req *GetFriendsRequest) isValid() error {
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

// GetFriendsList handles requests to retrieve the friends list for an email address
func (h Handler) GetFriendsList() http.HandlerFunc {
	return handler.ErrHandler(func(w http.ResponseWriter, r *http.Request) error {
		var req GetFriendsRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return handler.HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "Invalid JSON request",
			}
		}

		if err := req.isValid(); err != nil {
			return err
		}

		friendsList, count, err := h.relationshipService.GetFriendsList(r.Context(), relationship.GetFriendsInput{Email: req.Email})
		if err != nil {
			return handler.ConvertErr(err)
		}

		json.NewEncoder(w).Encode(handler.GetFriendsSuccess{
			Success: true,
			Friends: friendsList,
			Count:   count,
		})

		return nil
	})
}
