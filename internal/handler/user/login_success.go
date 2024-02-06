package user

// LoginSuccess is success reponse when user login successfully
type LoginSuccess struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}
