package server

import (
	"context"
	"encoding/json"
	"github.com/ray-laboratories/saturn/application"
	"github.com/ray-laboratories/saturn/types"
	"net/http"
	"strings"
)

type AuthHandler struct {
	auth application.AuthService
}

func NewAuthHandler(auth application.AuthService) *AuthHandler {
	return &AuthHandler{
		auth: auth,
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Unmarshal JSON
	var userRequest types.AccessorRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to decode request body",
		})
		return
	}

	// Attempt login
	token, err := ah.auth.Login(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Provide token
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
	return
}

func (ah *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Grab token
	token := r.Context().Value(TokenKey).(string)
	err := ah.auth.Logout(r.Context(), token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Unmarshal JSON
	var userRequest types.AccessorRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to decode request body",
		})
		return
	}

	// Register user
	token, err := ah.auth.Register(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Give back token
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
	return
}

func (ah *AuthHandler) Validate(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(TokenKey).(string)
	user, err := ah.auth.Validate(r.Context(), token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"token": token,
		"user":  user,
	})
	return
}

const (
	TokenKey = "token"
	UserKey  = "user"
)

func BearerToken(auth application.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Authorization header is missing",
				})
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Invalid authorization header format",
				})
				return
			}

			token := parts[1]

			user, err := auth.Validate(r.Context(), token)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error": err.Error(),
				})
				return
			}

			ctx := context.WithValue(r.Context(), TokenKey, token)
			ctx = context.WithValue(ctx, UserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
