package utils

import (
	"strconv"
	"strings"
)

// Limpia una entrada de texto (espacios y saltos de línea)
func SanitizeString(input string) string {
	return strings.TrimSpace(input)
}

// Conversión segura de string a int
func SafeAtoi(value string) (int, error) {
	value = strings.TrimSpace(value)
	return strconv.Atoi(value)
}
