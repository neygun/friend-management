package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrHandler handles errors from handler methods and responses JSON
func ErrHandler(handlerFunc func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := handlerFunc(w, r); err != nil {
			herr, ok := err.(HandlerErr)
			if ok {
				w.WriteHeader(herr.Code)

				json.NewEncoder(w).Encode(Response{
					Code:        herr.Code,
					Description: herr.Description,
				})
				log.Println(err.Error())

				return
			}

			w.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(w).Encode(Response{
				Code:        http.StatusInternalServerError,
				Description: "Internal Server Error",
			})
			log.Println(err.Error())
		}
	}
}
