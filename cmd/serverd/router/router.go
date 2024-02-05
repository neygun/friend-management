package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/neygun/friend-management/internal/handler/relationship"
	"github.com/neygun/friend-management/internal/handler/user"
	"github.com/neygun/friend-management/internal/service/authentication"
)

var tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

// Init defines API endpoints
func Init(r *chi.Mux, userHandler user.Handler, relationshipHandler relationship.Handler, authService authentication.Service) {

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens
		r.Use(AuthenticationMiddleware(authService))

		r.Post("/friends", relationshipHandler.CreateFriendConnection())
		r.Post("/friends/list", relationshipHandler.GetFriendsList())
		r.Post("/friends/common", relationshipHandler.GetCommonFriends())
		r.Post("/friends/subscription", relationshipHandler.CreateSubscription())
		r.Post("/friends/block", relationshipHandler.CreateBlock())
		r.Post("/friends/recipients", relationshipHandler.GetEmailsReceivingUpdates())
	})

	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens
		r.Use(AuthenticationMiddleware(authService))

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

func AuthenticationMiddleware(svc authentication.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())
			if err != nil {
				errorHandler(w, err, http.StatusUnauthorized)
				return
			}

			tokenString := jwtauth.TokenFromHeader(r)
			isBlacklisted, err := svc.CheckBlacklistedToken(r.Context(), tokenString)
			if err != nil {
				errorHandler(w, errInternalServerError, http.StatusInternalServerError)
				return
			}

			if token == nil || jwt.Validate(token) != nil || isBlacklisted {
				errorHandler(w, errUnauthorized, http.StatusUnauthorized)
				return
			}

			// Token is authenticated, pass it through
			next.ServeHTTP(w, r)
		})
	}
}

// func AuthenticationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token, _, err := jwtauth.FromContext(r.Context())
// 		if err != nil {
// 			errorHandler(w, err, http.StatusUnauthorized)
// 			return
// 		}

// 		tokenString := jwtauth.TokenFromHeader(r)
// 		isBlacklisted, err := isBlacklisted(tokenString)
// 		if err != nil {
// 			errorHandler(w, errInternalServerError, http.StatusInternalServerError)
// 			return
// 		}

// 		if token == nil || jwt.Validate(token) != nil || isBlacklisted {
// 			errorHandler(w, errUnauthorized, http.StatusUnauthorized)
// 			return
// 		}

// 		// Token is authenticated, pass it through
// 		next.ServeHTTP(w, r)
// 	})
// }
