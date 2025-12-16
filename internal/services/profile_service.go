package services

import (
	"errors"
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
)

type ProfileService struct {
	db *database.DB
}

func NewProfileService(db *database.DB) *ProfileService {
    return &ProfileService{db: db}
}

func (ps *ProfileService) CreateProfile(UserID int, name string, age int) error {
	if age < 0 {
	    return errors.New("Edad inválida")
	}
	if name == "" {
	    return errors.New("Nombre de perfil requerido")
	}
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	_, err = stmt.Exec(UserID, name, age)
	return err
}

func (ps *ProfileService) GetProfilesByUserID(userID int) ([]models.Profile, error) {
	rows, err := ps.db.Conn.Query("SELECT id, name, age, is_child FROM profiles WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []models.Profile
	for rows.Next() {
		var p models.Profile
		err := rows.Scan(&p.ID, &p.UserID, &p.Name, &p.Age)
		if err != nil {
		}
		profiles = append(profiles, p)
	}
	return profiles, nil
}