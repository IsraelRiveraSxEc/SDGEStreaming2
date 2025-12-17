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

	if !profile.IsAllowedFor(contentAllowed.MinAge) {
	    t.Error("El perfil tiene acceso, contenido permitido")
	}

	contentBlocked := models.Content{
	    MinAge: 16,
	}
	if profile.IsAllowedFor(contentBlocked.MinAge) {
	    t.Error("El perfil no tiene acceso, contenido bloqueado")	    
	}
}

func TestSubscriptionAccessLogic(t *testing.T) {
	subscription := models.UserSubscription{
	    Plan: models.PlanFree,
		Active: true,
	}

	if subscription.Plan != models.PlanFree {
	    t.Error("Plan free")
	}

	if !subscription.Active {
	    t.Error("Activo un plan")
	}
}