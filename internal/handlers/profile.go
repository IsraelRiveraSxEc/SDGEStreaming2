package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/services"	
)

// Handler para crear un perfil
func CreateProfileHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID int    `json:"user_id"`
			Name   string `json:"name"`
			Age    int    `json:"age"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		profileService := services.NewProfileService(db)
		if err := profileService.CreateProfile(req.UserID, req.Name, req.Age); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "perfil creado correctamente",
		})
	}
}

// Handler para listar perfiles por usuario
func GetProfilesHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userIDStr := params["user_id"]

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "user_id inválido", http.StatusBadRequest)
			return
		}

		profileService := services.NewProfileService(db)
		profiles, err := profileService.GetProfilesByUserID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(profiles)
	}
}