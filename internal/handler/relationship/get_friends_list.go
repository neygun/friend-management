package relationship

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/pkg/util"
)

func isValid(emailParam string) error {
	// check if the email is valid
	if !util.IsEmail(emailParam) {
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
		emailParam := chi.URLParam(r, "email")

		if err := isValid(emailParam); err != nil {
			return err
		}

		friendsList, count, err := h.relationshipService.GetFriendsList(r.Context(), emailParam)
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
