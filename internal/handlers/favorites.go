package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/services"
)

// Handler para agregar un favorito
func AddFavoriteHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ProfileID int `json:"profile_id"`
			ContentID int `json:"content_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		favoriteService := services.NewFavoriteService(db)
		if err := favoriteService.AddFavorite(req.ProfileID, req.ContentID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "contenido agregado a favoritos",
		})
	}
}

// Handler para listar favoritos de un perfil
func GetFavoritesHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		profileIDStr := params["profile_id"]

		profileID, err := strconv.Atoi(profileIDStr)
		if err != nil {
			http.Error(w, "profile_id inválido", http.StatusBadRequest)
			return
		}

		favoriteService := services.NewFavoriteService(db)
		favorites, err := favoriteService.GetFavorites(profileID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(favorites)
	}
}