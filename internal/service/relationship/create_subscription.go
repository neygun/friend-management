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
func (s service) CreateSubscription(ctx context.Context, input CreateSubscriptionInput) (model.Relationship, error) {
	// get users by emails
	users, err := s.userRepo.GetByCriteria(ctx, model.UserFilter{Emails: []string{input.Requestor, input.Target}})
	if err != nil {
		return model.Relationship{}, err
	}

	// check if there are 2 users with the emails
	if len(users) != 2 {
		return model.Relationship{}, ErrUserNotFound
	}

	requestorID, targetID := users[1].ID, users[0].ID
	if users[0].Email == input.Requestor {
		requestorID, targetID = users[0].ID, users[1].ID
	}

	// get relationships
	relationships, err := s.relationshipRepo.GetByCriteria(ctx, model.RelationshipFilter{
		RequestorID: requestorID,
		TargetID:    targetID})
	if err != nil {
		return model.Relationship{}, err
	}

	rels := make(map[model.RelationshipType]model.Relationship)
	for _, r := range relationships {
		rels[r.Type] = r
	}

	switch {
	// subscription exists
	case containsKey(rels, model.RelationshipTypeSubscribe):
		return model.Relationship{}, ErrSubscriptionExists
	// block exists
	case containsKey(rels, model.RelationshipTypeBlock):
		r := rels[model.RelationshipTypeBlock]
		r.Type = model.RelationshipTypeSubscribe
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
			Type:        model.RelationshipTypeSubscribe,
		}
		subscription, err = s.relationshipRepo.Create(ctx, subscription)
		if err != nil {
			return model.Relationship{}, err
		}
		return subscription, nil
	}

}

func containsKey[TKey model.RelationshipType, TModel any](m map[TKey]TModel, key TKey) bool {
	_, exists := m[key]
	return exists
}
