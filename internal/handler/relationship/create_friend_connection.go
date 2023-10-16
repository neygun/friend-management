package relationship

import (
	"encoding/json"
	"net/http"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/service/relationship"
	"github.com/neygun/friend-management/pkg/util"
)

// FriendConnectionRequest matches JSON request to create a friend connection
type FriendConnectionRequest struct {
	Friends []string
}

func (req FriendConnectionRequest) isValid() error {
	// check if number of emails = 2
	if len(req.Friends) != 2 {
		return handler.HandlerErr{
			Code:        http.StatusBadRequest,
			Description: "The number of emails must be 2",
		}
	}

	// check valid emails
	for _, v := range req.Friends {
		if !util.IsEmail(v) {
			return handler.HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "Invalid email",
			}
		}
	}
	return nil
}

// CreateFriendConnection handles requests to create friend connection
func (h Handler) CreateFriendConnection() http.HandlerFunc {
	return handler.ErrHandler(func(w http.ResponseWriter, r *http.Request) error {
		var req FriendConnectionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return handler.HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "Invalid JSON request",
			}
		}

		req.isValid()

		if _, err := h.relationshipService.CreateFriendConnection(r.Context(), relationship.FriendConnectionInput{
			Friends: req.Friends,
		}); err != nil {
			return handler.ConvertErr(err)
		}

		json.NewEncoder(w).Encode(model.SuccessResponse{
			Success: true,
		})

		return nil
	})
}
