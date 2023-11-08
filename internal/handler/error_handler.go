package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorHandler handles errors from handler methods and responses JSON
func ErrorHandler(handlerFunc func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Catch handler error
		if err := handlerFunc(w, r); err != nil {
			// Write HandlerError if err is HandlerError
			herr, ok := err.(HandlerError)
			if ok {
				w.WriteHeader(herr.Code)

				json.NewEncoder(w).Encode(Response{
					Code:        herr.Code,
					Description: herr.Description,
				})
				log.Println(err.Error())

				return
			}

			// Write internal error
			w.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(w).Encode(Response{
				Code:        http.StatusInternalServerError,
				Description: "Internal Server Error",
			})
			log.Println(err.Error())
		}
	}
}
