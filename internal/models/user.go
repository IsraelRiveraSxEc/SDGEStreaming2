package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Role int

const (
	RoleUser Role = iota
	RoleAdmin
	RoleDirector
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`			// no se expone en JSON
	Role Role `json:"role"`
	CreatedAt string `json:"created_at"`
}

func (u *User) SetPassword(password string) error {
	if len(password) < 8 {
		return errors.New("La contraseña debe tener al menos 8 caracteres")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}