package relationship

// GetFriendsSuccess is success reponse when retrieving the friends list for an email address
type GetFriendsSuccess struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}
