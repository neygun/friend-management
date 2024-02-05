package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
	"github.com/neygun/friend-management/cmd/serverd/router"
	"github.com/neygun/friend-management/internal/cache/authentication"
	"github.com/neygun/friend-management/internal/handler/relationship"
	"github.com/neygun/friend-management/internal/handler/user"
	relationshipRepository "github.com/neygun/friend-management/internal/repository/relationship"
	userRepository "github.com/neygun/friend-management/internal/repository/user"
	authService "github.com/neygun/friend-management/internal/service/authentication"
	relationshipService "github.com/neygun/friend-management/internal/service/relationship"
	userService "github.com/neygun/friend-management/internal/service/user"
	"github.com/neygun/friend-management/pkg/util"
)

func route(userHandler user.Handler, relationshipHandler relationship.Handler, authService authService.Service) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	router.Init(r, userHandler, relationshipHandler, authService)

	return r
}

func main() {
	db, err := util.InitDB()
	if err != nil {
		panic(err)
	}

	redisClient := util.NewRedisClient()

	relationshipRepo := relationshipRepository.New(db)
	userRepo := userRepository.New(db)
	authRepo := authentication.New(redisClient)

	userService := userService.New(userRepo, authRepo)
	userHandler := user.New(userService)
	relationshipService := relationshipService.New(userRepo, relationshipRepo)
	relationshipHandler := relationship.New(relationshipService)
	authService := authService.New(authRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	log.Println("Running on port " + port)
	http.ListenAndServe(":"+port, route(userHandler, relationshipHandler, authService))
}
