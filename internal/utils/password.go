package utils

import "golang.org/x/crypto/bcrypt"

package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Genera el hash de una contraseña
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Compara una contraseña en texto plano con su hash
func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
	return err == nil
}

