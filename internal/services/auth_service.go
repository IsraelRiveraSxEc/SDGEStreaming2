package services

import (
	"errors"
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
)

type AuthService struct {
	db *database.DB
}

func NewAuthService(db *database.DB) *AuthService {
    return &AuthService{db: db}
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
func (a *AuthService) Register (email, password string) error {
	user := &models.User{
		Email: email,
		Role: models.RoleUser,
	}
	if err := user.SetPassword(password); err != nil {
		return err
	}
	stmt, err := a.db.Conn.Prepare( "INSERT INTO users(email, password, is_admin) VALUES (?, ?, ?)",)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Email, user.Password, user.Role)
	return err
}

func (a *AuthService) Login(email, password string) (*models.User, error) {
	row := a.db.Conn.QueryRow(
		"SELECT id, email, password, is_admin FROM users WHERE email = ?", email,
	)
	
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role); err != nil {
		return nil, errors.New("Credenciales incorrectas")
	}
	
	if !user.CheckPassword(password) {
		return nil, errors.New("Credenciales incorrectas")
	}

	return &user, nil
}