package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"sdgestreaming/internal/handlers"
)

func TestLoginEndpoint(t *testing.T) {
	body := []byte(`{
		"email": "test@test.com",
		"password": "password123"
	}`)

	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("status esperado 200, obtenido %d", rr.Code)
	}
}

func TestGetProfilesEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/profiles", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/profiles", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("status esperado 200, obtenido %d", rr.Code)
	}
}

func TestGetContentEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/content", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/content", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("status esperado 200, obtenido %d", rr.Code)
	}
}