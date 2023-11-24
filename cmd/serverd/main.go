package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
	"github.com/neygun/friend-management/cmd/serverd/router"
	"github.com/neygun/friend-management/internal/handler/relationship"
	"github.com/neygun/friend-management/internal/handler/user"
	relationshipRepository "github.com/neygun/friend-management/internal/repository/relationship"
	userRepository "github.com/neygun/friend-management/internal/repository/user"
	relationshipService "github.com/neygun/friend-management/internal/service/relationship"
	userService "github.com/neygun/friend-management/internal/service/user"
	"github.com/neygun/friend-management/pkg/util"
)

func route(userHandler user.Handler, relationshipHandler relationship.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	router.Init(r, userHandler, relationshipHandler)

	return r
}

func main() {
	db, err := util.InitDB()
	if err != nil {
		panic(err)
	}

	relationshipRepo := relationshipRepository.New(db)
	userRepo := userRepository.New(db)
	bpe := userService.BCryptPasswordEncoder{}

	userService := userService.New(userRepo, bpe)
	userHandler := user.New(userService)

	relationshipService := relationshipService.New(userRepo, relationshipRepo)
	relationshipHandler := relationship.New(relationshipService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	log.Println("Running on port " + port)
	http.ListenAndServe(":"+port, route(userHandler, relationshipHandler))
}
