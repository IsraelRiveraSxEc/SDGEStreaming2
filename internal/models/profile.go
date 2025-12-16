package models

type AgeRating string

const (
	AgeTP  AgeRating = "TP"
	Age7   AgeRating = "7+"
	Age12  AgeRating = "12+"
	Age15  AgeRating = "15+"
	Age18  AgeRating = "18+"
)

type Profile struct {
	ID       int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Preferences map[string]string `json:"preferences"`
}

func (p *Profile) AgeRating() AgeRating {
	switch {
	case p.Age < 7:
		return AgeTP
	case p.Age < 12:
		return Age7
	case p.Age < 15:
		return Age12
	case p.Age < 18:
		return Age15
	default:
		return Age18
	}
}

func (p *Profile) IsAllowedFor(minAge int) bool {
	return p.Age >= minAge
}