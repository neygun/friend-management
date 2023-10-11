package relationship

import (
	"encoding/json"
	"net/http"

	"github.com/neygun/friend-management/internal/handler"
)

type FriendConnectionRequest struct {
	friends []string
}

func (h RelationshipHandler) CreateFriendConnection() http.HandlerFunc {
	return ErrHandler(func(w http.ResponseWriter, r *http.Request) error {
		var input FriendConnectionRequest
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "Invalid JSON request",
			}
		}

		// check valid emails
		for _, v := range input.friends {
			if !handler.IsEmail(v) {
				return HandlerErr{
					Code:        http.StatusBadRequest,
					Description: "Invalid email",
				}
			}
		}

		return nil
	})
}
