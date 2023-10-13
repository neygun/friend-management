package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/neygun/friend-management/internal/handler/relationship"
	"github.com/neygun/friend-management/internal/handler/user"
)

func InitRouter(r *chi.Mux, userHandler user.UserHandler, relationshipHandler relationship.RelationshipHandler) {
	r.Post("/user", userHandler.CreateUser())
	r.Post("/friend-connection", relationshipHandler.CreateFriendConnection())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

}
