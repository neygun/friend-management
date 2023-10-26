package relationship

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/neygun/friend-management/internal/handler"
	"github.com/neygun/friend-management/internal/service/relationship"
	"github.com/neygun/friend-management/pkg/util"
)

// CreateSubscriptionRequest matches JSON request to create a subscription
type CreateSubscriptionRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

func (req *CreateSubscriptionRequest) isValid() error {
	// trim space
	req.Requestor = strings.TrimSpace(req.Requestor)
	req.Target = strings.TrimSpace(req.Target)

	// check if requestor exists
	if req.Requestor == "" {
		return handler.HandlerErr{
			Code:        http.StatusBadRequest,
			Description: "Missing requestor field",
		}
	}

	// check if target exists
	if req.Target == "" {
		return handler.HandlerErr{
			Code:        http.StatusBadRequest,
			Description: "Missing target field",
		}
	}

	// check if the emails is valid
	if !util.IsEmail(req.Requestor) || !util.IsEmail(req.Target) {
		return handler.HandlerErr{
			Code:        http.StatusBadRequest,
			Description: "Invalid email",
		}
	}

	// check if requestor and target are the same
	if req.Requestor == req.Target {
		return handler.HandlerErr{
			Code:        http.StatusBadRequest,
			Description: "Requestor and target are the same",
		}
	}
	return nil
}

// CreateSubscription handles requests to create a subscription
func (h Handler) CreateSubscription() http.HandlerFunc {
	return handler.ErrHandler(func(w http.ResponseWriter, r *http.Request) error {
		var req CreateSubscriptionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return handler.HandlerErr{
				Code:        http.StatusBadRequest,
				Description: "Invalid JSON request",
			}
		}

		if err := req.isValid(); err != nil {
			return err
		}

		if _, err := h.relationshipService.CreateSubscription(r.Context(), relationship.CreateSubscriptionInput{
			Requestor: req.Requestor,
			Target:    req.Target,
		}); err != nil {
			return handler.ConvertErr(err)
		}

		json.NewEncoder(w).Encode(handler.SuccessResponse{
			Success: true,
		})

		return nil
	})
}
