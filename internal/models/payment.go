package models

import (
	"time"
)

type Payment struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	PlanID int `json:"plan_id"`
	Amount float64 `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Method string `json:"method"`
}