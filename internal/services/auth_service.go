package services

import (
	"database/sql"
	"errors"
	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
)

// Variables globales
var (
	ErrInvalidCredentials = errors.New("Credenciales invalidas")
	ErrEmailAlreadyExists = errors.New("Email ya registrado")
)

func NewAuthService(db *database.DB) *AuthService {
    return &AuthService{db: db}
}
// Registro
func (s *AuthService) Register(name, email, password, role string) (*models.User, error) {
	// Verificar si el email ya existe
	var exists bool
	err := s.db.Conn.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailAlreadyExists
	}

	// Crear el usuario
    user := &models.User{
        Name:     name,
        Email:    email,
        Password: password,
        Role:     role,
    }
	// Hashear contraseña
    if err = user.SetPassword(password); err != nil {
		return nil, err
	}
	return user, nil
}
// Login
func (s *AuthService) Login(email, password string) (*models.User, error) {
	var user models.User

	if	err := s.db.Conn.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdateAt)
		if err != nil {
			return nil, ErrInvalidCredentails
	}
	return nil & err
}
type AuthService struct {
	db *database.DB
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