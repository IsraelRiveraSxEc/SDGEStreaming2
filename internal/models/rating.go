package models

import "time"

type Rating struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	ContentID int `json:"content_id"`
	Value float64 `json:"value"` // Rango esperado de 0 a 10
	CreatedAt time.Time `json:"created_at"`
}