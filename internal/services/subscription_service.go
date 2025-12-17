package services

import (
	"errors"
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
)

type SubscriptionService struct {
	db *database.DB
}

func NewSubscriptionService(db *database.DB) *SubscriptionService {
	return &SubscriptionService{db: db}
}

// Asigna un plan de suscripción a un usuario
func (ss *SubscriptionService) AssignPlan(userID int, plan models.UserSubscription) error {
	stmt, err := ss.db.Conn.Prepare(
		"INSERT INTO subscriptions (user_id, plan_id, active) VALUES (?, ?, ?)",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, plan.ID, true)
	return err
}

// Valida si un usuario tiene acceso según su plan activo
func (ss *SubscriptionService) HasAccess(userID int, requiredPlan string) (bool, error) {
	row := ss.db.Conn.QueryRow(
		`SELECT sp.name
		 FROM subscriptions s
		 JOIN subscription_plans sp ON s.plan_id = sp.id
		 WHERE s.user_id = ? AND s.active = 1`,
		userID,
	)

	var planName string
	if err := row.Scan(&planName); err != nil {
		return false, errors.New("usuario sin suscripción activa")
	}

	switch requiredPlan {
	case "free":
		return true, nil
	case "standard":
		return planName == "standard" || planName == "premium", nil
	case "premium":
		return planName == "premium", nil
	default:
		return false, errors.New("plan requerido inválido")
	}
}