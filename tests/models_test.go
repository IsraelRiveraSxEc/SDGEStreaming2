package tests

import (
	"testing"
	"sdgestreaming/internal/models"
)

func TestUserPasswordHash(t *testing.T) {
	user := &models.User{}
	err := user.SetPassword("micontraseña123")
	if err != nil {
		t.Errorf("Error al hashear la contraseña: %v", err)
	}
	if !user.CheckPassword("micontraseña123") {
		t.Error("Las contraseñas no coinciden")
	}
}