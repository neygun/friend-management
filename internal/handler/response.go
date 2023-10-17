package handler

// Response represents JSON responses
type Response struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
