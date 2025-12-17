package models

type SubscriptionPlanType string

const (
	Free SubscriptionPlanType = "free"
	Standard SubscriptionPlanType = "standard"
	Premium SubscriptionPlanType = "premium"
)

type SubscriptionPlan struct {
	ID   int    `json:"id"`
	Plan SubscriptionPlanType `json:"plan"` // gratis, estandar o premium 
	Price float64 `json:"price"`
	MaxProfiles int `json:"max_profiles"`
	Features []string `json:"features"`
	IsActive bool `json:"is_active"`
}