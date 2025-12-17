package models

import (
	"fmt"
)

type Content struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"` // audiovisual o audio.
	Genre string `json:"genre"`
	Duration int `json:"duration"`
	Year int `json:"year"`
	Artist string `json:"artist"`
	Rating float64 `json:"rating"`
	MinAge int `json:min_age"`
}
// GetFormattedRating devuelve el rating con formato legible
func (c *Content) GetFormattedRating() string {
	if c.Rating == 10.0 {
		return "10"
	}
	return fmt.Sprintf("%.1f", c.Rating)
}