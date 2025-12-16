package services

import (
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
)

type CatalogService struct {
	db *database.DB
}

func NewCatalogService(db *database.DB) *CatalogService {
	return &CatalogService{db: db}
}

// Búsqueda general por título
func (cs *CatalogService) SearchByTitle(title string) ([]models.Content, error) {
	rows, err := cs.db.Conn.Query(
		`SELECT id, title, category, genre, duration, year, artist, rating, min_age
		 FROM contents
		 WHERE title LIKE ?`,
		"%"+title+"%",
	)
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

// Filtro por género
func (cs *CatalogService) FilterByGenre(genre string) ([]models.Content, error) {
	rows, err := cs.db.Conn.Query(
		`SELECT id, title, category, genre, duration, year, artist, rating, min_age
		 FROM contents
		 WHERE genre = ?`,
		genre,
	)
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