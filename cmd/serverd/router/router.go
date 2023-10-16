package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/neygun/friend-management/internal/handler/relationship"
	"github.com/neygun/friend-management/internal/handler/user"
)

// InitRouter init a router
func InitRouter(r *chi.Mux, userHandler user.Handler, relationshipHandler relationship.Handler) {
	r.Post("/user", userHandler.CreateUser())
	r.Post("/friend-connection", relationshipHandler.CreateFriendConnection())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

}
