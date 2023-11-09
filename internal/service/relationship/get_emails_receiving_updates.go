package relationship

import (
	"context"
	"regexp"

	"github.com/neygun/friend-management/internal/model"
)

// GetEmailsReceivingUpdatesInput is the input from request to retrieve emails that can receive updates from an email
type GetEmailsReceivingUpdatesInput struct {
	Sender string
	Text   string
}

// GetEmailsReceivingUpdates retrieves emails that can receive updates from an email
func (s service) GetEmailsReceivingUpdates(ctx context.Context, input GetEmailsReceivingUpdatesInput) ([]string, error) {
	// get sender by email
	sender, err := s.userRepo.GetByCriteria(ctx, model.UserFilter{Emails: []string{input.Sender}})
	if err != nil {
		return nil, err
	}

	// check if the sender exists
	if len(sender) == 0 {
		return nil, ErrUserNotFound
	}

	// extract emails mentioned in text
	pattern := regexp.MustCompile("[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*")
	mentionedEmails := pattern.FindAllString(input.Text, -1)

	var mentionedUserIDs []int64
	if len(mentionedEmails) != 0 {
		// get users by emails
		mentionedUsers, err := s.userRepo.GetByCriteria(ctx, model.UserFilter{Emails: mentionedEmails})
		if err != nil {
			return nil, err
		}
		// get user IDs
		mentionedUserIDs = make([]int64, len(mentionedUsers))
		for i, user := range mentionedUsers {
			mentionedUserIDs[i] = user.ID
		}
	}

	// get recipients
	recipients, err := s.relationshipRepo.GetEmailsReceivingUpdates(ctx, sender[0].ID, mentionedUserIDs)
	if err != nil {
		return nil, err
	}
	return recipients, nil
}
