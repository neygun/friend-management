package handler

// HandlerError represents response error for ErrorHandler function
type HandlerError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (e HandlerError) Error() string {
	return e.Description
}
