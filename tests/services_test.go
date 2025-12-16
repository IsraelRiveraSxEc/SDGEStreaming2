package tests

import (
	"testing"

	"sdgestreaming/internal/models"
)

func TestParentalControlLogic(t *testing.T) {
    profile := models.Profile{
        Age: 12,
    }

	contentAllowed := models.Content{
	    MinAge: 10,
	}

	if !profile.IsAllowedFor(contentAllowed) {
	    t.Error("El perfil tiene acceso al contenido permitido")
	}

	if profileIsnAllowedFor(contentBlocked.MinAge) {
	    t.Error("El perfil no tiene acceso, contenido bloqueado")	    
	}
}

func TestSubscriptionAccessLogic(t *testing.T) {
	subscription := models.Subscription{
	    Type: "free",
		Active: "true",
	}

	if subscription.Type != "free" {
	    t.Error("El tipo de suscripcion deberia ser free")
	}

	if !subscription.Active {
	    t.Error("La suscripcion deberia estar activa")
	}
}