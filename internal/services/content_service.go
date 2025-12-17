package services

import (
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
)

type ContentService struct{
	db *database.DB
}

func NewContentService(db *database.DB) *ContentService {
	return &ContentService{db: db}
}
// Obtiene todo el contenido permitido según la edad del perfil.
func (cs *ContentService) GetAllContentsForProfile(profile models.Profile) ([]models.Content, error) {
	rows, err := cs.db.Conn.Query(`SELECT id, title, type, genre, duration, year, artist, rating, min_age FROM contents WHERE min_age <= ?`,	profile.Age)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contents []models.Content
	for rows.Next() {
		var c models.Content
		if err := rows.Scan(
			&c.ID,
			&c.Title,
			&c.Category,
			&c.Genre,
			&c.Duration,
			&c.Year,
			&c.Artist,
			&c.Rating,
			&c.MinAge,
		); err != nil {
			return nil, err
		}
		contents = append(contents, c)
	}
	return contents, nil
}

func (cs *ContentService) SearchByTitleForProfile(profile models.Profile, title string) ([]models.Content, error) {
	rows, err := cs.db.Conn.Query("SELECT id, title, category, genre, duration, year, artist, rating, min_age FROM contents WHERE min_age <= ? AND title LIKE ?", profile.Age, "%"+title+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contents []models.Content
	for rows.Next() {
		var c models.Content
		if err := rows.Scan(
			&c.ID,
			&c.Title,
			&c.Category,
			&c.Genre,
			&c.Duration,
			&c.Year,
			&c.Artist,
			&c.Rating,
			&c.MinAge,
		); err != nil {
			return nil, err
		}
		
		contents = append(contents, c)
	}
	return contents, nil
}
