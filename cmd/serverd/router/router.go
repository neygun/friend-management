package router

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/neygun/friend-management/internal/handler/relationship"
	"github.com/neygun/friend-management/internal/handler/user"
	"github.com/neygun/friend-management/pkg/util"
	"github.com/redis/go-redis/v9"
)

var tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

// Init defines API endpoints
func Init(r *chi.Mux, userHandler user.Handler, relationshipHandler relationship.Handler) {

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens
		r.Use(AuthenticationMiddleware)

		r.Post("/friends", relationshipHandler.CreateFriendConnection())
		r.Post("/friends/list", relationshipHandler.GetFriendsList())
		r.Post("/friends/common", relationshipHandler.GetCommonFriends())
		r.Post("/friends/subscription", relationshipHandler.CreateSubscription())
		r.Post("/friends/block", relationshipHandler.CreateBlock())
		r.Post("/friends/recipients", relationshipHandler.GetEmailsReceivingUpdates())

		r.Post("/users/logout", userHandler.Logout())
	})

	r.Group(func(r chi.Router) {
		r.Post("/users", userHandler.CreateUser())
		r.Post("/users/login", userHandler.Login())

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("root."))
		})
	})
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		tokenString := jwtauth.TokenFromHeader(r)
		isBlacklisted, err := isBlacklisted(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if token == nil || jwt.Validate(token) != nil || isBlacklisted {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

func isBlacklisted(token string) (bool, error) {
	ctx := context.Background()
	client := util.NewRedisClient()
	result, err := client.Get(ctx, token).Result()
	if err == redis.Nil {
		// Token not found in the blacklist
		return false, nil
	} else if err != nil {
		return false, err
	}

	if result == "revoked" {
		return true, nil
	}

	return false, nil
}
