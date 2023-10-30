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

	requestorID, targetID := users[1].ID, users[0].ID
	if users[0].Email == createSubscriptionInput.Requestor {
		requestorID, targetID = users[0].ID, users[1].ID
	}

	// get relationships
	relationships, err := s.relationshipRepo.GetByCriteria(ctx, model.RelationshipFilter{
		RequestorID: requestorID,
		TargetID:    targetID})
	if err != nil {
		return model.Relationship{}, err
	}

	for _, r := range relationships {
		// if subscription already exists
		if r.Type == model.RelationshipTypeSubscribe.ToString() {
			return model.Relationship{}, ErrSubscriptionExists
		}

		// if block exists
		if r.Type == model.RelationshipTypeBlock.ToString() {
			r.Type = model.RelationshipTypeSubscribe.ToString()
			subscription, err := s.relationshipRepo.Update(ctx, r)
			if err != nil {
				return model.Relationship{}, err
			}
			return subscription, nil
		}
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
