package handlers

import (
	"encoding/json"
	"net/http"
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
	"sdgestreaming/internal/services"
)

// Handler de login
func LoginHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req services.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		authService := services.NewAuthService(db)
		user, err := authService.Login(req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// Handler de registro
func RegisterHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		authService := services.NewAuthService(db)
		if err := authService.Register(user.Email, user.Password, false); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "usuario registrado correctamente",
		})
	}
}