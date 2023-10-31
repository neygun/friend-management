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

	rels := make(map[string]model.Relationship)
	for _, r := range relationships {
		rels[r.Type] = r
	}

	switch {
	// subscription exists
	case containsKey(rels, model.RelationshipTypeSubscribe.ToString()):
		return model.Relationship{}, ErrSubscriptionExists
	// block exists
	case containsKey(rels, model.RelationshipTypeBlock.ToString()):
		r := rels[model.RelationshipTypeBlock.ToString()]
		r.Type = model.RelationshipTypeSubscribe.ToString()
		subscription, err := s.relationshipRepo.Update(ctx, r)
		if err != nil {
			return model.Relationship{}, err
		}
		return subscription, nil
	default:
		// create a new subscription
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

}

func containsKey(m map[string]model.Relationship, key string) bool {
	_, exists := m[key]
	return exists
}
