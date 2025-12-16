package models

type SubscriptionPlanType string

const (
	PlanFree SubscriptionPlanType = "free"
	PlanBasic SubscriptionPlanType = "standard"
	PlanPro SubscriptionPlanType = "premium"
)

type Subscription struct {
	ID   int    `json:"id"`
	Plan SubscriptionPlanType `json:"plan"` // gratis, estandar o premium 
	Price float64 `json:"price"`
	MaxProfiles int `json:"max_profiles"`
	Features []string `json:"features"`
	IsActive bool `json:"is_active"`
}