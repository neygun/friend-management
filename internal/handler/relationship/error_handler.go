package relationship

import "net/http"

type HandlerErr struct {
	Code        int
	Description string
}

func (e HandlerErr) Error() string {
	return e.Description
}

func ErrHandler(handlerFunc func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(w, r); err != nil {
			herr, ok := err.(HandlerErr)
			if ok {

			}
		}
	}
}
