package tests

import (
	"testing"
	"time"
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

func TestPaymentModel(t *testing.T) {
	payment := models.Payment{
		UserID:        1,
		Plan:          models.PlanPremium,
		Amount:        9.99,
		PaymentMethod: "card",
		PaidAt:        time.Now(),
	}

	if payment.Amount <= 0 {
		t.Error("El monto del pago debe ser mayor a 0")
	}

	if payment.Plan == "" {
		t.Error("El plan no debe estar vacío")
	}
}

func TestHistoryCompletion(t *testing.T) {
	h := models.History{
		Progress: 100,
	}

	if !h.IsCompleted() {
		t.Error("El contenido debería marcarse como completado")
	}
}

func TestFavoriteStruct(t *testing.T) {
	f := models.Favorite{
		ProfileID: 1,
		ContentID: 10,
	}

	if f.ProfileID == 0 || f.ContentID == 0 {
		t.Error("Favorite debe tener profile_id y content_id válidos")
	}
}

func TestRatingStruct(t *testing.T) {
	r := models.Rating{
		UserID:    1,
		ContentID: 5,
		Value:     9.5,
	}

	if r.Value < 0 || r.Value > 10 {
		t.Error("rating fuera de rango")
	}
}