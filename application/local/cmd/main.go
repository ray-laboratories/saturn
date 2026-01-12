package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ray-laboratories/saturn/application/local"
	"github.com/ray-laboratories/saturn/application/local/server"
	"net/http"
)

func main() {
	fmt.Println("Routing...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	fmt.Println("Configuring...")
	authSvc := local.NewDefaultAuthService("./app.db")
	authHandler := server.NewAuthHandler(authSvc)
	fmt.Println("Routing again...")
	r.With(server.BearerToken(authSvc)).Post("/logout", authHandler.Logout)
	r.Post("/login", authHandler.Login)
	r.Post("/register", authHandler.Register)
	r.With(server.BearerToken(authSvc)).Get("/validate", authHandler.Validate)
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":8365", r)
	if err != nil {
		panic(err)
	}
}
