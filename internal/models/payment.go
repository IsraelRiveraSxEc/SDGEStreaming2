package models

import "time"

type Payment struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Plan SubscriptionPlanType `json:"plan"`
	Amount float64 `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	PaidAt time.Time `json:"paid_at"`
}