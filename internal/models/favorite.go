package models

type Favorite struct {
	ID int `json:"id"`
	ProfileID int `json:"profile_id"`
	ContentID int `json:"content_id"`
	CreatedAt time.Time `json:"created_at"`
}