package models

type Rating struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	ContentID int `json:"content_id"`
	Value float64 `json:"value"`
}