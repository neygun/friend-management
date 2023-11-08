package relationship

// GetEmailsReceivingUpdatesSuccess is success reponse when retrieving emails that can receive updates from an email
type GetEmailsReceivingUpdatesSuccess struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}
