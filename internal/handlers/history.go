package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/services"
)

// Handler para obtener el historial de visualización de un perfil
func GetHistoryHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		profileIDStr := params["profile_id"]

		profileID, err := strconv.Atoi(profileIDStr)
		if err != nil {
			http.Error(w, "profile_id inválido", http.StatusBadRequest)
			return
		}

		historyService := services.NewHistoryService(db)
		history, err := historyService.GetHistoryByProfile(profileID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(history)
	}
}
