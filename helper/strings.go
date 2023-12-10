package helper

import "strings"

func ContainsUppercase(s string) bool {
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func ContainsDigit(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func ContainsSpecialChar(s string) bool {
	specialChars := "~`!@#$%^&*()-_+={}[]|;:'<>,.?/"
	for _, char := range s {
		if strings.ContainsRune(specialChars, char) {
			return true
		}
	}
	return false
}
