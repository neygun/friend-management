package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/neygun/friend-management/internal/handler/relationship"
	"github.com/neygun/friend-management/internal/handler/user"
)

// Init init a router
func Init(r *chi.Mux, userHandler user.Handler, relationshipHandler relationship.Handler) {
	r.Post("/users", userHandler.CreateUser())
	r.Post("/friends", relationshipHandler.CreateFriendConnection())
	r.Post("/friends/list", relationshipHandler.GetFriendsList())
	r.Post("/friends/common", relationshipHandler.GetCommonFriends())
	r.Post("/friends/subscription", relationshipHandler.CreateSubscription())
	r.Post("/friends/block", relationshipHandler.CreateBlock())
	r.Post("/friends/recipients", relationshipHandler.GetEmailsReceivingUpdates())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

}
