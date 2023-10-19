package handler

// HandlerErr represents errors that are not server errors
type HandlerErr struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// Error implemented
func (e HandlerErr) Error() string {
	return e.Description
}
