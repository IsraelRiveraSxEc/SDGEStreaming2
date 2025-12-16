package utils

import (
	"regexp"
	"strings"
)

// Valida formato de email
func IsValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	if email == "" {
		return false
	}

	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

// Valida números enteros dentro de un rango
func IsValidNumber(value int, min int, max int) bool {
	return value >= min && value <= max
}

// Valida texto no vacío y con longitud mínima
func IsValidText(text string, minLength int) bool {
	text = strings.TrimSpace(text)
	return len(text) >= minLength
}
