package models

type SubscriptionPlan string

const (
	PlanFree SubscriptionPlan = "free"
	PlanBasic SubscriptionPlan = "standard"
	PlanPro SubscriptionPlan = "premium"
)

type Subscription struct {
	ID   int    `json:"id"`
	Plan SubscriptionPlanType `json:"plan"` // gratis, estandar o premium 
	Price float64 `json:"price"`
	MaxProfiles int `json:"max_profiles"`
	Features []string `json:"features"`
	IsActive bool `json:"is_active"`
}