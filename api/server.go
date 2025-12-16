package api

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/handlers"
)

type Server struct {
	DB *database.DB
}

func NewServer(db *database.DB) *Server {
	return &Server{DB: db}
}

func (s *Server) Start(port string) error {
	router := mux.NewRouter()

	// Auth
	router.HandleFunc("/api/login", handlers.LoginHandler(s.DB)).Methods("POST")
	router.HandleFunc("/api/register", handlers.RegisterHandler(s.DB)).Methods("POST")

	// Profiles
	router.HandleFunc("/api/profiles", handlers.CreateProfileHandler(s.DB)).Methods("POST")
	router.HandleFunc("/api/profiles/{user_id}", handlers.GetProfilesHandler(s.DB)).Methods("GET")

	// Content
	router.HandleFunc("/api/content", handlers.GetAllContentHandler(s.DB)).Methods("GET")
	router.HandleFunc("/api/content/{id}", handlers.GetContentByIDHandler(s.DB)).Methods("GET")

	// Favorites
	router.HandleFunc("/api/favorites", handlers.AddFavoriteHandler(s.DB)).Methods("POST")
	router.HandleFunc("/api/favorites/{profile_id}", handlers.GetFavoritesHandler(s.DB)).Methods("GET")

	// History
	router.HandleFunc("/api/history/{profile_id}", handlers.GetHistoryHandler(s.DB)).Methods("GET")

	log.Println("Servidor iniciado en", port)
	return http.ListenAndServe(port, router)
}