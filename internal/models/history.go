package models

import (
	"time"
)

type History struct {
	ID int `json:"id"`
	ProfileID int `json:"profile_id"`
	ContentID int `json:"content_id"`
	Progress int `json:"progress"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
