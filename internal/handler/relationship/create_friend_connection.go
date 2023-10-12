package relationship

import (
	"encoding/json"
	"net/http"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/service/relationship"
)

type FriendConnectionRequest struct {
	Friends []string
}

func (h RelationshipHandler) CreateFriendConnection() http.HandlerFunc {
	return ErrHandler(func(w http.ResponseWriter, r *http.Request) error {
		var req FriendConnectionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "Invalid JSON request",
			}
		}

		if len(req.Friends) != 2 {
			return HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "The number of emails must be 2",
			}
		}

		// check valid emails
		for _, v := range req.Friends {
			if !handler.IsEmail(v) {
				return HandlerErr{
					Code:        http.StatusBadRequest,
					Description: "Invalid email",
				}
			}
		}

		if _, err := h.relationshipService.CreateFriendConnection(r.Context(), relationship.FriendConnectionInput{
			Friends: req.Friends,
		}); err != nil {
			return convertErr(err)
		}

		json.NewEncoder(w).Encode(model.Response{
			Code:        http.StatusOK,
			Description: "success",
		})

		return nil
	})
}
