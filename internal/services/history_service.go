package services

import (
	"time"
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
)

type HistoryService struct {
	db *database.DB
}

func NewHistoryService(db *database.DB) *HistoryService {
	return &HistoryService{db: db}
}

// Guarda o actualiza el progreso de reproducción de un contenido
func (hs *HistoryService) SaveProgress(profileID int, contentID int, progress int) error {
	stmt, err := hs.db.Conn.Prepare(
		`INSERT INTO history (profile_id, content_id, progress, updated_at)
		 VALUES (?, ?, ?, ?)
		 ON CONFLICT(profile_id, content_id)
		 DO UPDATE SET progress = ?, updated_at = ?`,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(
		profileID,
		contentID,
		progress,
		now,
		progress,
		now,
	)
	return err
}

// Lista el historial de visualización de un perfil
func (hs *HistoryService) GetHistoryByProfile(profileID int) ([]models.History, error) {
	rows, err := hs.db.Conn.Query(
		`SELECT id, profile_id, content_id, progress, updated_at
		 FROM history
		 WHERE profile_id = ?
		 ORDER BY updated_at DESC`,
		profileID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []models.History
	for rows.Next() {
		var h models.History
		if err := rows.Scan(
			&h.ID,
			&h.ProfileID,
			&h.ContentID,
			&h.Progress,
			&h.UpdatedAt,
		); err != nil {
			return nil, err
		}
		history = append(history, h)
	}

	return history, nil
}