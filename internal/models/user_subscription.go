package models

type UserSubscription struct {
    ID int `json:"id"`
    UserID int `json:"user_id"`
	Plan SubscriptionPlanType `json:"plan"`
	Active bool `json:"active"`
}