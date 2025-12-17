package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
	"sdgestreaming/internal/services"
)

// Handler para listar todo el contenido (sin control parental)
func GetAllContentHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentService := services.NewContentService(db)

		// Perfil simulado (ej. adulto)
		profile := models.Profile{Age: 100}

		contents, err := contentService.GetAllContentsForProfile(profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contents)
	}
}

// Handler para obtener contenido por ID
func GetContentByIDHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idStr := params["id"]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		row := db.Conn.QueryRow(
			`SELECT id, title, category, genre, duration, year, artist, rating, min_age
			 FROM contents WHERE id = ?`,
			id,
		)

		var c models.Content
		if err := row.Scan(
			&c.ID,
			&c.Title,
			&c.Category,
			&c.Genre,
			&c.Duration,
			&c.Year,
			&c.Artist,
			&c.Rating,
			&c.MinAge,
		); err != nil {
			http.Error(w, "contenido no encontrado", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c)
	}
}