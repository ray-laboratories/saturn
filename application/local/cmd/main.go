package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ray-laboratories/saturn/application/local"
	"github.com/ray-laboratories/saturn/application/local/server"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	authSvc := local.NewDefaultAuthService()
	authHandler := server.NewAuthHandler(authSvc)
	r.With(server.BearerToken(authSvc)).Post("/logout", authHandler.Logout)
	r.Post("/login", authHandler.Login)
	r.Post("/register", authHandler.Register)
	r.With(server.BearerToken(authSvc)).Get("/validate", authHandler.Validate)
	err := http.ListenAndServe(":8365", r)
	if err != nil {
		panic(err)
	}
}
