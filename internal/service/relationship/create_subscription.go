package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

// CreateSubscriptionInput is the input from request to create a subscription
type CreateSubscriptionInput struct {
	Requestor string
	Target    string
}

// CreateSubscription creates a subscription relationship
func (s service) CreateSubscription(ctx context.Context, createSubscriptionInput CreateSubscriptionInput) (model.Relationship, error) {
	// get users by emails
	users, err := s.userRepo.GetByCriteria(ctx, model.UserFilter{Emails: []string{createSubscriptionInput.Requestor, createSubscriptionInput.Target}})
	if err != nil {
		return model.Relationship{}, err
	}

	// check if there are 2 users with the emails
	if len(users) != 2 {
		return model.Relationship{}, ErrUserNotFound
	}

	var requestorID, targetID int64
	if users[0].Email == createSubscriptionInput.Requestor {
		requestorID = users[0].ID
		targetID = users[1].ID
	} else {
		requestorID = users[1].ID
		targetID = users[0].ID
	}

	// create subscription
	subscription := model.Relationship{
		RequestorID: requestorID,
		TargetID:    targetID,
		Type:        model.RelationshipTypeSubscribe.ToString(),
	}

	subscription, err = s.relationshipRepo.Create(ctx, subscription)
	if err != nil {
		return model.Relationship{}, err
	}

	return subscription, nil
}
