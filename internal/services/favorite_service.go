package services

import (
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
)

type FavoriteService struct {
	db *database.DB
}

func NewFavoriteService(db *database.DB) *FavoriteService {
	return &FavoriteService{db: db}
}

// Agrega un contenido a favoritos
func (fs *FavoriteService) AddFavorite(profileID int, contentID int) error {
	stmt, err := fs.db.Conn.Prepare(
		"INSERT INTO favorites (profile_id, content_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(profileID, contentID)
	return err
}

// Elimina un contenido de favoritos
func (fs *FavoriteService) RemoveFavorite(profileID int, contentID int) error {
	stmt, err := fs.db.Conn.Prepare(
		"DELETE FROM favorites WHERE profile_id = ? AND content_id = ?",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(profileID, contentID)
	return err
}

// Lista los favoritos de un perfil
func (fs *FavoriteService) GetFavorites(profileID int) ([]models.Content, error) {
	rows, err := fs.db.Conn.Query(
		`SELECT c.id, c.title, c.category, c.genre, c.duration, c.year, c.artist, c.rating, c.min_age
		 FROM favorites f
		 JOIN contents c ON f.content_id = c.id
		 WHERE f.profile_id = ?`,
		profileID,
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