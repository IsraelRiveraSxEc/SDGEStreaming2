package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	RoleUser Role = "user"
	RoleAdmin Role = "admin"
	RoleDirector Role = "director"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`			// no se expone en JSON
	Role Role `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
// Genera el hash bcrypt
func (u *User) SetPassword(plain string) error {
	if len(plain) < 8 {
		return errors.New("La contraseña debe tener al menos 8 caracteres")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	u.Password = string(hash)
	return nil
}
// Compara el hash con el texto plano
func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}