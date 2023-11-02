package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

// CreateBlockInput is the input from request to create a blocking relationship
type CreateBlockInput struct {
	Requestor string
	Target    string
}

// CreateBlock creates a blocking relationship
func (s service) CreateBlock(ctx context.Context, createBlockInput CreateBlockInput) (model.Relationship, error) {
	// get users by emails
	users, err := s.userRepo.GetByCriteria(ctx, model.UserFilter{Emails: []string{createBlockInput.Requestor, createBlockInput.Target}})
	if err != nil {
		return model.Relationship{}, err
	}

	// check if there are 2 users with the emails
	if len(users) != 2 {
		return model.Relationship{}, ErrUserNotFound
	}

	requestorID, targetID := users[1].ID, users[0].ID
	if users[0].Email == createBlockInput.Requestor {
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
	// block exists
	case containsKey(rels, model.RelationshipTypeBlock):
		return model.Relationship{}, ErrBlockExists
	// subscription exists
	case containsKey(rels, model.RelationshipTypeSubscribe):
		r := rels[model.RelationshipTypeSubscribe]
		r.Type = model.RelationshipTypeBlock
		block, err := s.relationshipRepo.Update(ctx, r)
		if err != nil {
			return model.Relationship{}, err
		}
		return block, nil
	default:
		// create a new blocking relationship
		block := model.Relationship{
			RequestorID: requestorID,
			TargetID:    targetID,
			Type:        model.RelationshipTypeBlock,
		}
		block, err = s.relationshipRepo.Create(ctx, block)
		if err != nil {
			return model.Relationship{}, err
		}
		return block, nil
	}

}
