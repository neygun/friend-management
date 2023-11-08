package relationship

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/service/relationship"
	"github.com/neygun/friend-management/pkg/util"
)

// GetEmailsReceivingUpdatesRequest matches JSON request to retrieve emails that can receive updates from an email
type GetEmailsReceivingUpdatesRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

func (req *GetEmailsReceivingUpdatesRequest) isValid() error {
	// trim space
	req.Sender = strings.TrimSpace(req.Sender)

	// check if sender exists
	if req.Sender == "" {
		return handler.HandlerError{
			Code:        http.StatusBadRequest,
			Description: "Missing sender field",
		}
	}

	// check if text exists
	if req.Text == "" {
		return handler.HandlerError{
			Code:        http.StatusBadRequest,
			Description: "Missing text field",
		}
	}

	// check if the sender's email is valid
	if !util.IsEmail(req.Sender) {
		return handler.HandlerError{
			Code:        http.StatusBadRequest,
			Description: "Sender's email is invalid",
		}
	}

	return nil
}

// GetEmailsReceivingUpdates handles requests to retrieve emails that can receive updates from an email
func (h Handler) GetEmailsReceivingUpdates() http.HandlerFunc {
	return handler.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		var req GetEmailsReceivingUpdatesRequest
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

		// Get emails receiving updates
		recipients, err := h.relationshipService.GetEmailsReceivingUpdates(r.Context(), relationship.GetEmailsReceivingUpdatesInput{
			Sender: req.Sender,
			Text:   req.Text,
		})
		if err != nil {
			return handler.ConvertError(err)
		}

		// Success
		json.NewEncoder(w).Encode(GetEmailsReceivingUpdatesSuccess{
			Success:    true,
			Recipients: recipients,
		})

		return nil
	})
}
